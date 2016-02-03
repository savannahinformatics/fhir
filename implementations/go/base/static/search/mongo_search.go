package search

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/intervention-engine/fhir/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoSearcher implements FHIR searches using the Mongo database.
type MongoSearcher struct {
	db *mgo.Database
}

// NewMongoSearcher creates a new instance of a MongoSearcher, given a pointer
// to an mgo.Database.
func NewMongoSearcher(db *mgo.Database) *MongoSearcher {
	return &MongoSearcher{db}
}

// CreateQuery takes a FHIR-based Query and returns a pointer to the
// corresponding mgo.Query.  The returned mgo.Query will obey any options
// passed in through the query string (such as _count and _offset) and will
// also use default options when none are passed in (e.g., count = 100).
// The caller is responsible for executing the returned query (allowing
// additional flexibility in how results are returned).
//
// CreateQuery CANNOT be used when the _include and _revinclude options
// are used (since CreateQuery can't support joins).
func (m *MongoSearcher) CreateQuery(query Query) *mgo.Query {
	return m.createQuery(query, true)
}

// CreateQueryWithoutOptions takes a FHIR-based Query and returns a pointer to
// the corresponding mgo.Query.  Any options passed in through the query (such
// as _count and _offset) are ignored and no default options are applied (e.g.,
// there is no set count / limit)  The caller is responsible for executing
// the returned query (allowing flexibility in how results are returned).
func (m *MongoSearcher) CreateQueryWithoutOptions(query Query) *mgo.Query {
	return m.createQuery(query, false)
}

func (m *MongoSearcher) createQuery(query Query, withOptions bool) *mgo.Query {
	c := m.db.C(models.PluralizeLowerResourceName(query.Resource))
	q := m.createQueryObject(query)
	mgoQuery := c.Find(q)

	// Horrible, horrible hack (for now) to ensure patients are sorted by name.  This is needed by
	// the frontend, else paging won't work correctly.  This should be removed when the general
	// sorting feature is implemented.
	if query.Resource == "Patient" {
		// To add insult to injury, mongo will not let us sort by family *and* given name:
		// Executor error: BadValue cannot sort with keys that are parallel arrays
		mgoQuery = mgoQuery.Sort("name.0.family.0" /*", name.0.given.0"*/, "_id")
	}

	if withOptions {
		o := query.Options()
		if o.Offset > 0 {
			mgoQuery = mgoQuery.Skip(o.Offset)
		}
		mgoQuery = mgoQuery.Limit(o.Count)
	}
	return mgoQuery
}

// CreatePipeline takes a FHIR-based Query and returns a pointer to the
// corresponding mgo.Pipe.  The returned mgo.Pipe will obey any options
// passed in through the query string (such as _count and _offset) and will
// also use default options when none are passed in (e.g., count = 100).
// The caller is responsible for executing the returned pipe (allowing
// additional flexibility in how results are returned).
//
// CreatePipeline must be used when the _include and _revinclude options
// are used (since CreateQuery can't support joins).
func (m *MongoSearcher) CreatePipeline(query Query) *mgo.Pipe {
	c := m.db.C(models.PluralizeLowerResourceName(query.Resource))
	p := []bson.M{{"$match": m.createQueryObject(query)}}

	// Horrible, horrible hack (for now) to ensure patients are sorted by name.  This is needed by
	// the frontend, else paging won't work correctly.  This should be removed when the general
	// sorting feature is implemented.
	if query.Resource == "Patient" {
		// To add insult to injury, mongo will not let us sort by family *and* given name:
		// Executor error: BadValue cannot sort with keys that are parallel arrays
		p = append(p, bson.M{"$sort": bson.M{"name.0.family.0": 1, "_id": 1}})
	}

	o := query.Options()
	// support for _offset
	if o.Offset > 0 {
		p = append(p, bson.M{"$skip": o.Offset})
	}
	// support for _count
	p = append(p, bson.M{"$limit": o.Count})

	// support for _include
	if len(o.Include) > 0 {
		for _, incl := range o.Include {
			for _, inclPath := range incl.Parameter.Paths {
				if inclPath.Type != "Reference" {
					continue
				}
				// Mongo paths shouldn't have the array indicators, so remove them
				localField := strings.Replace(inclPath.Path, "[]", "", -1) + ".referenceid"
				for i, inclTarget := range incl.Parameter.Targets {
					if inclTarget == "Any" {
						continue
					}
					from := models.PluralizeLowerResourceName(inclTarget)
					as := fmt.Sprintf("_included%sResourcesReferencedBy%s", inclTarget, strings.Title(incl.Parameter.Name))
					// If there are multiple paths, we need to store each path separately
					if len(incl.Parameter.Paths) > 1 {
						as += fmt.Sprintf("Path%d", i+1)
					}

					p = append(p, bson.M{"$lookup": bson.M{
						"from":         from,
						"localField":   localField,
						"foreignField": "_id",
						"as":           as,
					}})
				}
			}
		}
	}

	// support for _revinclude
	if len(o.RevInclude) > 0 {
		for _, incl := range o.RevInclude {
			// we only want parameters that have the search resource as their target
			targetsSearchResource := false
			for _, inclTarget := range incl.Parameter.Targets {
				if inclTarget == query.Resource || inclTarget == "Any" {
					targetsSearchResource = true
					break
				}
			}
			if !targetsSearchResource {
				continue
			}
			// it comes from the other resource collection
			from := models.PluralizeLowerResourceName(incl.Parameter.Resource)
			// iterate through the paths, adding a join to the pipeline for each one
			for i, inclPath := range incl.Parameter.Paths {
				if inclPath.Type != "Reference" {
					continue
				}
				// Mongo paths shouldn't have the array indicators, so remove them
				foreignField := strings.Replace(inclPath.Path, "[]", "", -1) + ".referenceid"
				as := fmt.Sprintf("_revIncluded%sResourcesReferencing%s", incl.Parameter.Resource, strings.Title(incl.Parameter.Name))
				// If there are multiple paths, we need to store each path separately
				if len(incl.Parameter.Paths) > 1 {
					as += fmt.Sprintf("Path%d", i+1)
				}

				p = append(p, bson.M{"$lookup": bson.M{
					"from":         from,
					"localField":   "_id",
					"foreignField": foreignField,
					"as":           as,
				}})

			}
		}
	}

	return c.Pipe(p)
}

func (m *MongoSearcher) createQueryObject(query Query) bson.M {
	result := bson.M{}
	for _, p := range m.createParamObjects(query.Resource, query.Params()) {
		merge(result, p)
	}
	return result
}

func (m *MongoSearcher) createParamObjects(resource string, params []SearchParam) []bson.M {
	results := make([]bson.M, len(params))
	for i, p := range params {
		panicOnUnsupportedFeatures(p)
		switch p := p.(type) {
		case *CompositeParam:
			results[i] = m.createCompositeQueryObject(resource, p)
		case *DateParam:
			results[i] = m.createDateQueryObject(p)
		case *NumberParam:
			results[i] = m.createNumberQueryObject(p)
		case *QuantityParam:
			results[i] = m.createQuantityQueryObject(p)
		case *ReferenceParam:
			results[i] = m.createReferenceQueryObject(p)
		case *StringParam:
			results[i] = m.createStringQueryObject(p)
		case *TokenParam:
			results[i] = m.createTokenQueryObject(p)
		case *URIParam:
			results[i] = m.createURIQueryObject(p)
		case *OrParam:
			results[i] = m.createOrQueryObject(resource, p)
		default:
			panic(createInternalServerError("MSG_PARAM_UNKNOWN", fmt.Sprintf("Parameter \"%s\" not understood", p)))
		}
	}

	return results
}

func panicOnUnsupportedFeatures(p SearchParam) {
	// No prefixes are supported except EQ (the default) and date prefixes
	_, isDate := p.(*DateParam)
	prefix := p.getInfo().Prefix
	if prefix != "" && prefix != EQ && !isDate {
		panic(createUnsupportedSearchError("MSG_PARAM_INVALID", fmt.Sprintf("Parameter \"%s\" content is invalid", p.getInfo().Name)))
	}

	// No modifiers are supported except for resource types in reference parameters
	_, isRef := p.(*ReferenceParam)
	modifier := p.getInfo().Modifier
	if modifier != "" {
		if _, ok := SearchParameterDictionary[modifier]; !isRef || !ok {
			panic(createUnsupportedSearchError("MSG_PARAM_MODIFIER_INVALID", fmt.Sprintf("Parameter \"%s\" modifier is invalid", p.getInfo().Name)))
		}
	}
}

func (m *MongoSearcher) createCompositeQueryObject(resource string, c *CompositeParam) bson.M {
	panic(createUnsupportedSearchError("MSG_PARAM_UNKNOWN", fmt.Sprintf("Parameter \"%s\" not understood", c.Name)))
}

func (m *MongoSearcher) createDateQueryObject(d *DateParam) bson.M {
	single := func(p SearchParamPath) bson.M {
		switch p.Type {
		case "date", "dateTime", "instant":
			return buildBSON(p.Path, dateSelector(d))
		case "Period":
			return buildBSON(p.Path, periodSelector(d))
		case "Timing":
			return buildBSON(p.Path+".event", dateSelector(d))
		default:
			return bson.M{}
		}
	}

	return orPaths(single, d.Paths)
}

// Note that this solution is not 100% correct because we don't represent dates as ranges in the
// database -- but the FHIR spec calls for this sort of behavior to correctly implement these
// searches.  An easy example is that while 2012-01-01 should be compared as the range from
// 00:00:00.000 to 23:59:59.999, we currently only compare against 00:00:00.000 -- so some things
// that should match, might not.
// TODO: Fix this via more complex search criteria or by a different representation in the database.
func dateSelector(d *DateParam) bson.M {
	var timeCriteria bson.M
	switch d.Prefix {
	case EQ:
		timeCriteria = bson.M{
			"$gte": d.Date.RangeLowIncl(),
			"$lt":  d.Date.RangeHighExcl(),
		}
	case GT:
		timeCriteria = bson.M{
			"$gt": d.Date.RangeLowIncl(),
		}
	case LT:
		timeCriteria = bson.M{
			"$lt": d.Date.RangeLowIncl(),
		}
	case GE:
		timeCriteria = bson.M{
			"$gte": d.Date.RangeLowIncl(),
		}
	case LE:
		timeCriteria = bson.M{
			"$lt": d.Date.RangeHighExcl(),
		}
	default:
		panic(createUnsupportedSearchError("MSG_PARAM_INVALID", fmt.Sprintf("Parameter \"%s\" content is invalid", d.Name)))
	}

	return bson.M{"time": timeCriteria}
}

// Note that this solution is not 100% correct because we don't represent dates as ranges in the
// database -- but the FHIR spec calls for this sort of behavior to correctly implement these
// searches.  An easy example is that while 2012-01-01 should be compared as the range from
// 00:00:00.000 to 23:59:59.999, we currently only compare against 00:00:00.000 -- so some things
// that should match, might not.
// TODO: Fix this via more complex search criteria or by a different representation in the database.
func periodSelector(d *DateParam) bson.M {
	switch d.Prefix {
	case EQ:
		return bson.M{
			"start.time": bson.M{
				"$gte": d.Date.RangeLowIncl(),
			},
			"end.time": bson.M{
				"$lt": d.Date.RangeHighExcl(),
			},
		}
	case GT:
		return bson.M{
			"$or": []bson.M{
				bson.M{
					"end.time": bson.M{
						"$gt": d.Date.RangeLowIncl(),
					},
				},
				// Also support instances where period exists, but end is null (ongoing)
				bson.M{
					"$ne": nil,
					"end": nil,
				},
			},
		}
	case LT:
		return bson.M{
			"$or": []bson.M{
				bson.M{
					"start.time": bson.M{
						"$lt": d.Date.RangeLowIncl(),
					},
				},
				// Also support instances where period exists, but start is null
				bson.M{
					"$ne":   nil,
					"start": nil,
				},
			},
		}
	case GE:
		return bson.M{
			"$or": []bson.M{
				bson.M{
					"start.time": bson.M{
						"$gte": d.Date.RangeLowIncl(),
					},
				},
				bson.M{
					"end.time": bson.M{
						"$gt": d.Date.RangeLowIncl(),
					},
				},
				// Also support instances where period exists, but end is null (ongoing)
				bson.M{
					"$ne": nil,
					"end": nil,
				},
			},
		}
	case LE:
		return bson.M{
			"$or": []bson.M{
				bson.M{
					"end.time": bson.M{
						"$lt": d.Date.RangeHighExcl(),
					},
				},
				bson.M{
					"start.time": bson.M{
						"$lt": d.Date.RangeLowIncl(),
					},
				},
				// Also support instances where period exists, but start is null
				bson.M{
					"$ne":   nil,
					"start": nil,
				},
			},
		}
	}
	panic(createUnsupportedSearchError("MSG_PARAM_INVALID", fmt.Sprintf("Parameter \"%s\" content is invalid", d.Name)))
}

func (m *MongoSearcher) createNumberQueryObject(n *NumberParam) bson.M {
	single := func(p SearchParamPath) bson.M {
		l, _ := n.Number.RangeLowIncl().Float64()
		h, _ := n.Number.RangeHighExcl().Float64()
		return buildBSON(p.Path, bson.M{
			"$gte": l,
			"$lt":  h,
		})
	}

	return orPaths(single, n.Paths)
}

func (m *MongoSearcher) createQuantityQueryObject(q *QuantityParam) bson.M {
	single := func(p SearchParamPath) bson.M {
		l, _ := q.Number.RangeLowIncl().Float64()
		h, _ := q.Number.RangeHighExcl().Float64()
		criteria := bson.M{
			"value": bson.M{
				"$gte": l,
				"$lt":  h,
			},
		}
		if q.System == "" {
			criteria["$or"] = []bson.M{
				bson.M{"code": ci(q.Code)},
				bson.M{"unit": ci(q.Code)},
			}
		} else {
			criteria["code"] = ci(q.Code)
			criteria["system"] = ci(q.System)
		}
		return buildBSON(p.Path, criteria)
	}

	return orPaths(single, q.Paths)
}

func (m *MongoSearcher) createReferenceQueryObject(r *ReferenceParam) bson.M {
	single := func(p SearchParamPath) bson.M {
		if p.Type == "Resource" {
			return m.createInlinedReferenceQueryObject(r, p)
		}
		criteria := bson.M{}
		switch ref := r.Reference.(type) {
		case LocalReference:
			criteria["referenceid"] = ci(ref.ID)
			if ref.Type != "" {
				criteria["type"] = ref.Type
			}
		case ExternalReference:
			criteria["reference"] = ci(ref.URL)
		case ChainedQueryReference:
			// Since MongoDB does not support cross-collection searches, we must break this into two:
			// (1) perform search against referenced collection using chained search Query
			// (2) use ID results from first query to build second query
			// TODO: Investigate if new Mongo 3.2 $lookup pipeline feature might be an improvement
			var idObjs []struct {
				ID string `bson:"_id"`
			}
			q := m.CreateQueryWithoutOptions(ref.ChainedQuery)
			q.Select(bson.M{"_id": 1}).All(&idObjs)
			ids := make([]string, len(idObjs))
			for i := range idObjs {
				ids[i] = idObjs[i].ID
			}
			criteria["referenceid"] = bson.M{"$in": ids}
			if ref.Type != "" {
				criteria["type"] = ref.Type
			}
		}
		return buildBSON(p.Path, criteria)
	}

	return orPaths(single, r.Paths)
}

func (m *MongoSearcher) createInlinedReferenceQueryObject(r *ReferenceParam, p SearchParamPath) bson.M {
	criteria := bson.M{}
	switch ref := r.Reference.(type) {
	case LocalReference:
		if ref.Type != "" {
			criteria["resourceType"] = ref.Type
		}
		criteria["_id"] = ci(ref.ID)
	case ChainedQueryReference:
		criteria = m.createQueryObject(ref.ChainedQuery)
		if ref.Type != "" {
			criteria["resourceType"] = ref.Type
		}
	case ExternalReference:
		panic(createUnsupportedSearchError("MSG_PARAM_INVALID", fmt.Sprintf("Parameter \"%s\" content is invalid", r.Name)))
	}
	return buildBSON(p.Path, criteria)
}

func (m *MongoSearcher) createStringQueryObject(s *StringParam) bson.M {
	single := func(p SearchParamPath) bson.M {
		switch p.Type {
		case "HumanName":
			return buildBSON(p.Path, bson.M{
				"$or": []bson.M{
					bson.M{"text": cisw(s.String)},
					bson.M{"family": cisw(s.String)},
					bson.M{"given": cisw(s.String)},
				},
			})
		case "Address":
			return buildBSON(p.Path, bson.M{
				"$or": []bson.M{
					bson.M{"text": cisw(s.String)},
					bson.M{"line": cisw(s.String)},
					bson.M{"city": cisw(s.String)},
					bson.M{"state": cisw(s.String)},
					bson.M{"postalCode": cisw(s.String)},
					bson.M{"country": cisw(s.String)},
				},
			})
		default:
			var criteria bson.RegEx
			if s.Name == "_id" {
				criteria = ci(s.String)
			} else {
				criteria = cisw(s.String)
			}
			return buildBSON(p.Path, criteria)
		}
	}

	return orPaths(single, s.Paths)
}

func (m *MongoSearcher) createTokenQueryObject(t *TokenParam) bson.M {
	single := func(p SearchParamPath) bson.M {
		criteria := bson.M{}
		switch p.Type {
		case "Coding":
			criteria = bson.M{}
			criteria["code"] = ci(t.Code)
			if !t.AnySystem {
				criteria["system"] = ci(t.System)
			}
		case "CodeableConcept":
			if t.AnySystem {
				criteria["coding.code"] = ci(t.Code)
			} else {
				criteria["coding"] = bson.M{"$elemMatch": bson.M{"system": ci(t.System), "code": ci(t.Code)}}
			}
		case "Identifier":
			criteria["value"] = ci(t.Code)
			if !t.AnySystem {
				criteria["system"] = ci(t.System)
			}
		case "ContactPoint":
			criteria["value"] = ci(t.Code)
			if !t.AnySystem {
				criteria["use"] = ci(t.System)
			}
		case "code", "boolean", "string", "id":
			// criteria isn't a bson, so just return the right answer
			return buildBSON(p.Path, ci(t.Code))
		}

		return buildBSON(p.Path, criteria)
	}

	return orPaths(single, t.Paths)
}

func (m *MongoSearcher) createURIQueryObject(u *URIParam) bson.M {
	single := func(p SearchParamPath) bson.M {
		return buildBSON(p.Path, u.URI)
	}

	return orPaths(single, u.Paths)
}

func (m *MongoSearcher) createOrQueryObject(resource string, o *OrParam) bson.M {
	return bson.M{
		"$or": m.createParamObjects(resource, o.Items),
	}
}

func createOpOutcome(severity, code, detailsCode, detailsDisplay string) *models.OperationOutcome {
	outcome := &models.OperationOutcome{
		Issue: []models.OperationOutcomeIssueComponent{
			models.OperationOutcomeIssueComponent{
				Severity: severity,
				Code:     code,
			},
		},
	}

	if detailsCode != "" {
		outcome.Issue[0].Details = &models.CodeableConcept{
			Coding: []models.Coding{
				models.Coding{
					Code:    detailsCode,
					System:  "http://hl7.org/fhir/ValueSet/operation-outcome",
					Display: detailsDisplay},
			},
			Text: detailsDisplay,
		}
	}

	return outcome
}

// Error is an interface for search errors, providing an HTTP status and operation outcome
type Error struct {
	HTTPStatus       int
	OperationOutcome *models.OperationOutcome
}

func createUnsupportedSearchError(code, display string) *Error {
	return &Error{
		HTTPStatus:       http.StatusNotImplemented,
		OperationOutcome: createOpOutcome("error", "not-supported", code, display),
	}
}

func createInvalidSearchError(code, display string) *Error {
	return &Error{
		HTTPStatus:       http.StatusBadRequest,
		OperationOutcome: createOpOutcome("error", "processing", code, display),
	}
}

func createInternalServerError(code, display string) *Error {
	return &Error{
		HTTPStatus:       http.StatusInternalServerError,
		OperationOutcome: createOpOutcome("fatal", "exception", code, display),
	}
}

func buildBSON(path string, criteria interface{}) bson.M {
	result := bson.M{}

	// First fix the indexers so "[0]entry.resource" becomes "entry.0.resource"
	re := regexp.MustCompile("\\[(\\d+)\\]([^\\.]+)")
	indexedPath := re.ReplaceAllString(path, "$2.$1")
	normalizedPath := strings.Replace(indexedPath, "[]", "", -1)
	bCriteria, ok := criteria.(bson.M)
	if ok {
		pathRegex := regexp.MustCompile("(.*\\[\\][^\\.]*)\\.?([^\\[\\]]*)")
		if m := pathRegex.FindStringSubmatch(indexedPath); m != nil && len(bCriteria) > 1 {
			// Need to use an $elemMatch because there is an array in the path
			// and the search criteria is a composite
			left := strings.Replace(m[1], "[]", "", -1)
			right := m[2]
			var resultCriteria bson.M
			if len(right) > 0 {
				resultCriteria = bson.M{}
				for k, v := range bCriteria {
					// Pull out the $or and process it separately under $elemMatch
					if isQueryOperator(k) {
						processQueryOperatorCriteria(right, k, v, resultCriteria)
					} else {
						resultCriteria[fmt.Sprintf("%s.%s", right, k)] = v
					}
				}
			} else {
				resultCriteria = bCriteria
			}
			result[left] = bson.M{"$elemMatch": resultCriteria}
		} else {
			// Path has no array or criteria is singular
			for k, v := range bCriteria {
				// Pull out the $or and process it separately as top level condition
				if isQueryOperator(k) {
					processQueryOperatorCriteria(indexedPath, k, v, result)
				} else {
					result[fmt.Sprintf("%s.%s", normalizedPath, k)] = v
				}
			}
		}
	} else {
		// Criteria is singular, so we don't care about arrays
		result[normalizedPath] = criteria
	}

	return result
}

func isQueryOperator(key string) bool {
	return len(key) > 0 && key[0] == '$'
}

func processQueryOperatorCriteria(path string, key string, value interface{}, result bson.M) {
	switch key {
	case "$or":
		processOrCriteria(path, value, result)
	default:
		criteria, ok := result[path]
		if !ok {
			criteria = bson.M{}
			result[path] = criteria
		}
		criteria.(bson.M)[key] = value
	}
}

func processOrCriteria(path string, orValue interface{}, result bson.M) {
	if ors, ok := orValue.([]bson.M); ok {
		newOrs := make([]bson.M, len(ors))
		for i := range ors {
			newOrs[i] = buildBSON(path, ors[i])
		}
		result["$or"] = newOrs
	} else {
		panic(createInternalServerError("", ""))
	}
}

// Case-insensitive match
func ci(s string) bson.RegEx {
	return bson.RegEx{Pattern: fmt.Sprintf("^%s$", regexp.QuoteMeta(s)), Options: "i"}
}

// Case-insensitive starts-with
func cisw(s string) bson.RegEx {
	return bson.RegEx{Pattern: fmt.Sprintf("^%s", regexp.QuoteMeta(s)), Options: "i"}
}

// When multiple paths are present, they should be represented as an OR.
// objFunc is a function that generates a single query for a path
func orPaths(objFunc func(SearchParamPath) bson.M, paths []SearchParamPath) bson.M {
	results := make([]bson.M, 0, len(paths))
	for i := range paths {
		result := objFunc(paths[i])
		// If the bson is just an $or, then bring the components up to the top-level $or
		if len(result) == 1 && result["$or"] != nil {
			nestedOrs := result["$or"].([]bson.M)
			for j := range nestedOrs {
				results = append(results, nestedOrs[j])
			}
		} else {
			results = append(results, objFunc(paths[i]))
		}
	}

	if len(results) == 1 {
		return results[0]
	}

	return bson.M{"$or": results}
}

func merge(into bson.M, from bson.M) {
	var and []bson.M
	if intoAnd, ok := into["$and"]; ok {
		and = intoAnd.([]bson.M)
	}

	for k, v := range from {
		if k == "$and" {
			and = append(and, v.([]bson.M)...)
		} else if _, ok := into[k]; ok {
			and = append(and, bson.M{k: v})
		} else {
			into[k] = v
		}
	}

	if len(and) > 0 {
		into["$and"] = and
	}
}
