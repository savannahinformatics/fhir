// Copyright (c) 2011-2017, HL7, Inc & The MITRE Corporation
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted provided that the following conditions are met:
//
//     * Redistributions of source code must retain the above copyright notice, this
//       list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above copyright notice,
//       this list of conditions and the following disclaimer in the documentation
//       and/or other materials provided with the distribution.
//     * Neither the name of HL7 nor the names of its contributors may be used to
//       endorse or promote products derived from this software without specific
//       prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
// IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT,
// INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
// NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
// PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
// WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package models

import (
    "encoding/json"
    "fmt"
    "errors"
)

type ResearchDefinition struct {
    DomainResource `bson:",inline"`
    Url string `bson:"url,omitempty" json:"url,omitempty"`
    Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
    Version string `bson:"version,omitempty" json:"version,omitempty"`
    Name string `bson:"name,omitempty" json:"name,omitempty"`
    Title string `bson:"title,omitempty" json:"title,omitempty"`
    ShortTitle string `bson:"shortTitle,omitempty" json:"shortTitle,omitempty"`
    Subtitle string `bson:"subtitle,omitempty" json:"subtitle,omitempty"`
    Status string `bson:"status,omitempty" json:"status,omitempty"`
    Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty"`
    SubjectCodeableConcept *CodeableConcept `bson:"subjectCodeableConcept,omitempty" json:"subjectCodeableConcept,omitempty"`
    SubjectReference *Reference `bson:"subjectReference,omitempty" json:"subjectReference,omitempty"`
    Date *FHIRDateTime `bson:"date,omitempty" json:"date,omitempty"`
    Publisher string `bson:"publisher,omitempty" json:"publisher,omitempty"`
    Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
    Description string `bson:"description,omitempty" json:"description,omitempty"`
    Comment []string `bson:"comment,omitempty" json:"comment,omitempty"`
    UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
    Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
    Purpose string `bson:"purpose,omitempty" json:"purpose,omitempty"`
    Usage string `bson:"usage,omitempty" json:"usage,omitempty"`
    Copyright string `bson:"copyright,omitempty" json:"copyright,omitempty"`
    ApprovalDate *FHIRDateTime `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
    LastReviewDate *FHIRDateTime `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
    EffectivePeriod *Period `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
    Topic []CodeableConcept `bson:"topic,omitempty" json:"topic,omitempty"`
    Author []ContactDetail `bson:"author,omitempty" json:"author,omitempty"`
    Editor []ContactDetail `bson:"editor,omitempty" json:"editor,omitempty"`
    Reviewer []ContactDetail `bson:"reviewer,omitempty" json:"reviewer,omitempty"`
    Endorser []ContactDetail `bson:"endorser,omitempty" json:"endorser,omitempty"`
    RelatedArtifact []RelatedArtifact `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
    Library []canonical `bson:"library,omitempty" json:"library,omitempty"`
    Population *Reference `bson:"population,omitempty" json:"population,omitempty"`
    Exposure *Reference `bson:"exposure,omitempty" json:"exposure,omitempty"`
    ExposureAlternative *Reference `bson:"exposureAlternative,omitempty" json:"exposureAlternative,omitempty"`
    Outcome *Reference `bson:"outcome,omitempty" json:"outcome,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (resource *ResearchDefinition) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "ResearchDefinition"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*resource)
}



// "researchDefinition" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type researchDefinition ResearchDefinition

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *ResearchDefinition) UnmarshalJSON(data []byte) (err error) {
	x2 := researchDefinition{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
    			x2.Contained[i] = MapToResource(x2.Contained[i], true)
    		}
		}
		*x = ResearchDefinition(x2)
		return x.checkResourceType()
	}
	return
}

func (x *ResearchDefinition) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "ResearchDefinition"
	} else if x.ResourceType != "ResearchDefinition" {
		return errors.New(fmt.Sprintf("Expected resourceType to be ResearchDefinition, instead received %s", x.ResourceType))
	}
	return nil
}



type ResearchDefinitionPlus struct {
    ResearchDefinition `bson:",inline"`
    ResearchDefinitionPlusRelatedResources `bson:",inline"`
}

type ResearchDefinitionPlusRelatedResources struct {
    IncludedLibraryResourcesReferencedByDependsonPath1 *[]Library `bson:"_includedLibraryResourcesReferencedByDependsonPath1,omitempty"`
    IncludedLibraryResourcesReferencedByDependsonPath2 *[]Library `bson:"_includedLibraryResourcesReferencedByDependsonPath2,omitempty"`
    RevIncludedAppointmentResourcesReferencingSupportinginfo *[]Appointment `bson:"_revIncludedAppointmentResourcesReferencingSupportinginfo,omitempty"`
    RevIncludedEventDefinitionResourcesReferencingSuccessor *[]EventDefinition `bson:"_revIncludedEventDefinitionResourcesReferencingSuccessor,omitempty"`
    RevIncludedEventDefinitionResourcesReferencingDerivedfrom *[]EventDefinition `bson:"_revIncludedEventDefinitionResourcesReferencingDerivedfrom,omitempty"`
    RevIncludedEventDefinitionResourcesReferencingPredecessor *[]EventDefinition `bson:"_revIncludedEventDefinitionResourcesReferencingPredecessor,omitempty"`
    RevIncludedEventDefinitionResourcesReferencingComposedof *[]EventDefinition `bson:"_revIncludedEventDefinitionResourcesReferencingComposedof,omitempty"`
    RevIncludedEventDefinitionResourcesReferencingDependson *[]EventDefinition `bson:"_revIncludedEventDefinitionResourcesReferencingDependson,omitempty"`
    RevIncludedDocumentManifestResourcesReferencingItem *[]DocumentManifest `bson:"_revIncludedDocumentManifestResourcesReferencingItem,omitempty"`
    RevIncludedDocumentManifestResourcesReferencingRelatedref *[]DocumentManifest `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
    RevIncludedConsentResourcesReferencingData *[]Consent `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
    RevIncludedMeasureResourcesReferencingSuccessor *[]Measure `bson:"_revIncludedMeasureResourcesReferencingSuccessor,omitempty"`
    RevIncludedMeasureResourcesReferencingDerivedfrom *[]Measure `bson:"_revIncludedMeasureResourcesReferencingDerivedfrom,omitempty"`
    RevIncludedMeasureResourcesReferencingPredecessor *[]Measure `bson:"_revIncludedMeasureResourcesReferencingPredecessor,omitempty"`
    RevIncludedMeasureResourcesReferencingComposedof *[]Measure `bson:"_revIncludedMeasureResourcesReferencingComposedof,omitempty"`
    RevIncludedMeasureResourcesReferencingDependsonPath1 *[]Measure `bson:"_revIncludedMeasureResourcesReferencingDependsonPath1,omitempty"`
    RevIncludedMeasureResourcesReferencingDependsonPath2 *[]Measure `bson:"_revIncludedMeasureResourcesReferencingDependsonPath2,omitempty"`
    RevIncludedDocumentReferenceResourcesReferencingRelated *[]DocumentReference `bson:"_revIncludedDocumentReferenceResourcesReferencingRelated,omitempty"`
    RevIncludedMeasureReportResourcesReferencingEvaluatedresource *[]MeasureReport `bson:"_revIncludedMeasureReportResourcesReferencingEvaluatedresource,omitempty"`
    RevIncludedVerificationResultResourcesReferencingTarget *[]VerificationResult `bson:"_revIncludedVerificationResultResourcesReferencingTarget,omitempty"`
    RevIncludedContractResourcesReferencingSubject *[]Contract `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
    RevIncludedPaymentNoticeResourcesReferencingRequest *[]PaymentNotice `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
    RevIncludedPaymentNoticeResourcesReferencingResponse *[]PaymentNotice `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
    RevIncludedResearchDefinitionResourcesReferencingSuccessor *[]ResearchDefinition `bson:"_revIncludedResearchDefinitionResourcesReferencingSuccessor,omitempty"`
    RevIncludedResearchDefinitionResourcesReferencingDerivedfrom *[]ResearchDefinition `bson:"_revIncludedResearchDefinitionResourcesReferencingDerivedfrom,omitempty"`
    RevIncludedResearchDefinitionResourcesReferencingPredecessor *[]ResearchDefinition `bson:"_revIncludedResearchDefinitionResourcesReferencingPredecessor,omitempty"`
    RevIncludedResearchDefinitionResourcesReferencingComposedof *[]ResearchDefinition `bson:"_revIncludedResearchDefinitionResourcesReferencingComposedof,omitempty"`
    RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 *[]ResearchDefinition `bson:"_revIncludedResearchDefinitionResourcesReferencingDependsonPath1,omitempty"`
    RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 *[]ResearchDefinition `bson:"_revIncludedResearchDefinitionResourcesReferencingDependsonPath2,omitempty"`
    RevIncludedImplementationGuideResourcesReferencingResource *[]ImplementationGuide `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
    RevIncludedResearchElementDefinitionResourcesReferencingSuccessor *[]ResearchElementDefinition `bson:"_revIncludedResearchElementDefinitionResourcesReferencingSuccessor,omitempty"`
    RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom *[]ResearchElementDefinition `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDerivedfrom,omitempty"`
    RevIncludedResearchElementDefinitionResourcesReferencingPredecessor *[]ResearchElementDefinition `bson:"_revIncludedResearchElementDefinitionResourcesReferencingPredecessor,omitempty"`
    RevIncludedResearchElementDefinitionResourcesReferencingComposedof *[]ResearchElementDefinition `bson:"_revIncludedResearchElementDefinitionResourcesReferencingComposedof,omitempty"`
    RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 *[]ResearchElementDefinition `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDependsonPath1,omitempty"`
    RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 *[]ResearchElementDefinition `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDependsonPath2,omitempty"`
    RevIncludedCommunicationResourcesReferencingPartof *[]Communication `bson:"_revIncludedCommunicationResourcesReferencingPartof,omitempty"`
    RevIncludedCommunicationResourcesReferencingBasedon *[]Communication `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
    RevIncludedActivityDefinitionResourcesReferencingSuccessor *[]ActivityDefinition `bson:"_revIncludedActivityDefinitionResourcesReferencingSuccessor,omitempty"`
    RevIncludedActivityDefinitionResourcesReferencingDerivedfrom *[]ActivityDefinition `bson:"_revIncludedActivityDefinitionResourcesReferencingDerivedfrom,omitempty"`
    RevIncludedActivityDefinitionResourcesReferencingPredecessor *[]ActivityDefinition `bson:"_revIncludedActivityDefinitionResourcesReferencingPredecessor,omitempty"`
    RevIncludedActivityDefinitionResourcesReferencingComposedof *[]ActivityDefinition `bson:"_revIncludedActivityDefinitionResourcesReferencingComposedof,omitempty"`
    RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 *[]ActivityDefinition `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath1,omitempty"`
    RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 *[]ActivityDefinition `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath2,omitempty"`
    RevIncludedLinkageResourcesReferencingItem *[]Linkage `bson:"_revIncludedLinkageResourcesReferencingItem,omitempty"`
    RevIncludedLinkageResourcesReferencingSource *[]Linkage `bson:"_revIncludedLinkageResourcesReferencingSource,omitempty"`
    RevIncludedDeviceRequestResourcesReferencingBasedon *[]DeviceRequest `bson:"_revIncludedDeviceRequestResourcesReferencingBasedon,omitempty"`
    RevIncludedDeviceRequestResourcesReferencingPriorrequest *[]DeviceRequest `bson:"_revIncludedDeviceRequestResourcesReferencingPriorrequest,omitempty"`
    RevIncludedMessageHeaderResourcesReferencingFocus *[]MessageHeader `bson:"_revIncludedMessageHeaderResourcesReferencingFocus,omitempty"`
    RevIncludedImmunizationRecommendationResourcesReferencingInformation *[]ImmunizationRecommendation `bson:"_revIncludedImmunizationRecommendationResourcesReferencingInformation,omitempty"`
    RevIncludedProvenanceResourcesReferencingEntity *[]Provenance `bson:"_revIncludedProvenanceResourcesReferencingEntity,omitempty"`
    RevIncludedProvenanceResourcesReferencingTarget *[]Provenance `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
    RevIncludedTaskResourcesReferencingSubject *[]Task `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
    RevIncludedTaskResourcesReferencingFocus *[]Task `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
    RevIncludedTaskResourcesReferencingBasedon *[]Task `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
    RevIncludedListResourcesReferencingItem *[]List `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
    RevIncludedEvidenceVariableResourcesReferencingSuccessor *[]EvidenceVariable `bson:"_revIncludedEvidenceVariableResourcesReferencingSuccessor,omitempty"`
    RevIncludedEvidenceVariableResourcesReferencingDerivedfrom *[]EvidenceVariable `bson:"_revIncludedEvidenceVariableResourcesReferencingDerivedfrom,omitempty"`
    RevIncludedEvidenceVariableResourcesReferencingPredecessor *[]EvidenceVariable `bson:"_revIncludedEvidenceVariableResourcesReferencingPredecessor,omitempty"`
    RevIncludedEvidenceVariableResourcesReferencingComposedof *[]EvidenceVariable `bson:"_revIncludedEvidenceVariableResourcesReferencingComposedof,omitempty"`
    RevIncludedEvidenceVariableResourcesReferencingDependson *[]EvidenceVariable `bson:"_revIncludedEvidenceVariableResourcesReferencingDependson,omitempty"`
    RevIncludedObservationResourcesReferencingFocus *[]Observation `bson:"_revIncludedObservationResourcesReferencingFocus,omitempty"`
    RevIncludedLibraryResourcesReferencingSuccessor *[]Library `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
    RevIncludedLibraryResourcesReferencingDerivedfrom *[]Library `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
    RevIncludedLibraryResourcesReferencingPredecessor *[]Library `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
    RevIncludedLibraryResourcesReferencingComposedof *[]Library `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
    RevIncludedLibraryResourcesReferencingDependson *[]Library `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
    RevIncludedCommunicationRequestResourcesReferencingBasedon *[]CommunicationRequest `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
    RevIncludedBasicResourcesReferencingSubject *[]Basic `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
    RevIncludedEvidenceResourcesReferencingSuccessor *[]Evidence `bson:"_revIncludedEvidenceResourcesReferencingSuccessor,omitempty"`
    RevIncludedEvidenceResourcesReferencingDerivedfrom *[]Evidence `bson:"_revIncludedEvidenceResourcesReferencingDerivedfrom,omitempty"`
    RevIncludedEvidenceResourcesReferencingPredecessor *[]Evidence `bson:"_revIncludedEvidenceResourcesReferencingPredecessor,omitempty"`
    RevIncludedEvidenceResourcesReferencingComposedof *[]Evidence `bson:"_revIncludedEvidenceResourcesReferencingComposedof,omitempty"`
    RevIncludedEvidenceResourcesReferencingDependson *[]Evidence `bson:"_revIncludedEvidenceResourcesReferencingDependson,omitempty"`
    RevIncludedAuditEventResourcesReferencingEntity *[]AuditEvent `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
    RevIncludedConditionResourcesReferencingEvidencedetail *[]Condition `bson:"_revIncludedConditionResourcesReferencingEvidencedetail,omitempty"`
    RevIncludedCompositionResourcesReferencingSubject *[]Composition `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
    RevIncludedCompositionResourcesReferencingEntry *[]Composition `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
    RevIncludedDetectedIssueResourcesReferencingImplicated *[]DetectedIssue `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
    RevIncludedQuestionnaireResponseResourcesReferencingSubject *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
    RevIncludedClinicalImpressionResourcesReferencingSupportinginfo *[]ClinicalImpression `bson:"_revIncludedClinicalImpressionResourcesReferencingSupportinginfo,omitempty"`
    RevIncludedPlanDefinitionResourcesReferencingSuccessor *[]PlanDefinition `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
    RevIncludedPlanDefinitionResourcesReferencingDerivedfrom *[]PlanDefinition `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
    RevIncludedPlanDefinitionResourcesReferencingPredecessor *[]PlanDefinition `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
    RevIncludedPlanDefinitionResourcesReferencingComposedof *[]PlanDefinition `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
    RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 *[]PlanDefinition `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
    RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 *[]PlanDefinition `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
}

func (r *ResearchDefinitionPlusRelatedResources) GetIncludedLibraryResourceReferencedByDependsonPath1() (library *Library, err error) {
    if r.IncludedLibraryResourcesReferencedByDependsonPath1 == nil {
        err = errors.New("Included libraries not requested")
    } else if len(*r.IncludedLibraryResourcesReferencedByDependsonPath1) > 1 {
        err = fmt.Errorf("Expected 0 or 1 library, but found %d", len(*r.IncludedLibraryResourcesReferencedByDependsonPath1))
    } else if len(*r.IncludedLibraryResourcesReferencedByDependsonPath1) == 1 {
        library = &(*r.IncludedLibraryResourcesReferencedByDependsonPath1)[0]
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetIncludedLibraryResourcesReferencedByDependsonPath2() (libraries []Library, err error) {
    if r.IncludedLibraryResourcesReferencedByDependsonPath2 == nil {
        err = errors.New("Included libraries not requested")
    } else {
        libraries = *r.IncludedLibraryResourcesReferencedByDependsonPath2
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
    if r.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
        err = errors.New("RevIncluded appointments not requested")
    } else {
        appointments = *r.RevIncludedAppointmentResourcesReferencingSupportinginfo
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
    if r.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
        err = errors.New("RevIncluded eventDefinitions not requested")
    } else {
        eventDefinitions = *r.RevIncludedEventDefinitionResourcesReferencingSuccessor
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
    if r.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
        err = errors.New("RevIncluded eventDefinitions not requested")
    } else {
        eventDefinitions = *r.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
    if r.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
        err = errors.New("RevIncluded eventDefinitions not requested")
    } else {
        eventDefinitions = *r.RevIncludedEventDefinitionResourcesReferencingPredecessor
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
    if r.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
        err = errors.New("RevIncluded eventDefinitions not requested")
    } else {
        eventDefinitions = *r.RevIncludedEventDefinitionResourcesReferencingComposedof
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
    if r.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
        err = errors.New("RevIncluded eventDefinitions not requested")
    } else {
        eventDefinitions = *r.RevIncludedEventDefinitionResourcesReferencingDependson
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
    if r.RevIncludedDocumentManifestResourcesReferencingItem == nil {
        err = errors.New("RevIncluded documentManifests not requested")
    } else {
        documentManifests = *r.RevIncludedDocumentManifestResourcesReferencingItem
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
    if r.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
        err = errors.New("RevIncluded documentManifests not requested")
    } else {
        documentManifests = *r.RevIncludedDocumentManifestResourcesReferencingRelatedref
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
    if r.RevIncludedConsentResourcesReferencingData == nil {
        err = errors.New("RevIncluded consents not requested")
    } else {
        consents = *r.RevIncludedConsentResourcesReferencingData
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
    if r.RevIncludedMeasureResourcesReferencingSuccessor == nil {
        err = errors.New("RevIncluded measures not requested")
    } else {
        measures = *r.RevIncludedMeasureResourcesReferencingSuccessor
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
    if r.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
        err = errors.New("RevIncluded measures not requested")
    } else {
        measures = *r.RevIncludedMeasureResourcesReferencingDerivedfrom
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
    if r.RevIncludedMeasureResourcesReferencingPredecessor == nil {
        err = errors.New("RevIncluded measures not requested")
    } else {
        measures = *r.RevIncludedMeasureResourcesReferencingPredecessor
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
    if r.RevIncludedMeasureResourcesReferencingComposedof == nil {
        err = errors.New("RevIncluded measures not requested")
    } else {
        measures = *r.RevIncludedMeasureResourcesReferencingComposedof
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
    if r.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
        err = errors.New("RevIncluded measures not requested")
    } else {
        measures = *r.RevIncludedMeasureResourcesReferencingDependsonPath1
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
    if r.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
        err = errors.New("RevIncluded measures not requested")
    } else {
        measures = *r.RevIncludedMeasureResourcesReferencingDependsonPath2
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
    if r.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
        err = errors.New("RevIncluded documentReferences not requested")
    } else {
        documentReferences = *r.RevIncludedDocumentReferenceResourcesReferencingRelated
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
    if r.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
        err = errors.New("RevIncluded measureReports not requested")
    } else {
        measureReports = *r.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
    if r.RevIncludedVerificationResultResourcesReferencingTarget == nil {
        err = errors.New("RevIncluded verificationResults not requested")
    } else {
        verificationResults = *r.RevIncludedVerificationResultResourcesReferencingTarget
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
    if r.RevIncludedContractResourcesReferencingSubject == nil {
        err = errors.New("RevIncluded contracts not requested")
    } else {
        contracts = *r.RevIncludedContractResourcesReferencingSubject
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
    if r.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
        err = errors.New("RevIncluded paymentNotices not requested")
    } else {
        paymentNotices = *r.RevIncludedPaymentNoticeResourcesReferencingRequest
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
    if r.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
        err = errors.New("RevIncluded paymentNotices not requested")
    } else {
        paymentNotices = *r.RevIncludedPaymentNoticeResourcesReferencingResponse
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
    if r.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
        err = errors.New("RevIncluded researchDefinitions not requested")
    } else {
        researchDefinitions = *r.RevIncludedResearchDefinitionResourcesReferencingSuccessor
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
    if r.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
        err = errors.New("RevIncluded researchDefinitions not requested")
    } else {
        researchDefinitions = *r.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
    if r.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
        err = errors.New("RevIncluded researchDefinitions not requested")
    } else {
        researchDefinitions = *r.RevIncludedResearchDefinitionResourcesReferencingPredecessor
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
    if r.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
        err = errors.New("RevIncluded researchDefinitions not requested")
    } else {
        researchDefinitions = *r.RevIncludedResearchDefinitionResourcesReferencingComposedof
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
    if r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
        err = errors.New("RevIncluded researchDefinitions not requested")
    } else {
        researchDefinitions = *r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
    if r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
        err = errors.New("RevIncluded researchDefinitions not requested")
    } else {
        researchDefinitions = *r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
    if r.RevIncludedImplementationGuideResourcesReferencingResource == nil {
        err = errors.New("RevIncluded implementationGuides not requested")
    } else {
        implementationGuides = *r.RevIncludedImplementationGuideResourcesReferencingResource
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
    if r.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
        err = errors.New("RevIncluded researchElementDefinitions not requested")
    } else {
        researchElementDefinitions = *r.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
    if r.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
        err = errors.New("RevIncluded researchElementDefinitions not requested")
    } else {
        researchElementDefinitions = *r.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
    if r.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
        err = errors.New("RevIncluded researchElementDefinitions not requested")
    } else {
        researchElementDefinitions = *r.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
    if r.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
        err = errors.New("RevIncluded researchElementDefinitions not requested")
    } else {
        researchElementDefinitions = *r.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
    if r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
        err = errors.New("RevIncluded researchElementDefinitions not requested")
    } else {
        researchElementDefinitions = *r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
    if r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
        err = errors.New("RevIncluded researchElementDefinitions not requested")
    } else {
        researchElementDefinitions = *r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
    if r.RevIncludedCommunicationResourcesReferencingPartof == nil {
        err = errors.New("RevIncluded communications not requested")
    } else {
        communications = *r.RevIncludedCommunicationResourcesReferencingPartof
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
    if r.RevIncludedCommunicationResourcesReferencingBasedon == nil {
        err = errors.New("RevIncluded communications not requested")
    } else {
        communications = *r.RevIncludedCommunicationResourcesReferencingBasedon
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
    if r.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
        err = errors.New("RevIncluded activityDefinitions not requested")
    } else {
        activityDefinitions = *r.RevIncludedActivityDefinitionResourcesReferencingSuccessor
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
    if r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
        err = errors.New("RevIncluded activityDefinitions not requested")
    } else {
        activityDefinitions = *r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
    if r.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
        err = errors.New("RevIncluded activityDefinitions not requested")
    } else {
        activityDefinitions = *r.RevIncludedActivityDefinitionResourcesReferencingPredecessor
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
    if r.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
        err = errors.New("RevIncluded activityDefinitions not requested")
    } else {
        activityDefinitions = *r.RevIncludedActivityDefinitionResourcesReferencingComposedof
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
    if r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
        err = errors.New("RevIncluded activityDefinitions not requested")
    } else {
        activityDefinitions = *r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
    if r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
        err = errors.New("RevIncluded activityDefinitions not requested")
    } else {
        activityDefinitions = *r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
    if r.RevIncludedLinkageResourcesReferencingItem == nil {
        err = errors.New("RevIncluded linkages not requested")
    } else {
        linkages = *r.RevIncludedLinkageResourcesReferencingItem
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
    if r.RevIncludedLinkageResourcesReferencingSource == nil {
        err = errors.New("RevIncluded linkages not requested")
    } else {
        linkages = *r.RevIncludedLinkageResourcesReferencingSource
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
    if r.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
        err = errors.New("RevIncluded deviceRequests not requested")
    } else {
        deviceRequests = *r.RevIncludedDeviceRequestResourcesReferencingBasedon
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
    if r.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
        err = errors.New("RevIncluded deviceRequests not requested")
    } else {
        deviceRequests = *r.RevIncludedDeviceRequestResourcesReferencingPriorrequest
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
    if r.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
        err = errors.New("RevIncluded messageHeaders not requested")
    } else {
        messageHeaders = *r.RevIncludedMessageHeaderResourcesReferencingFocus
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
    if r.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
        err = errors.New("RevIncluded immunizationRecommendations not requested")
    } else {
        immunizationRecommendations = *r.RevIncludedImmunizationRecommendationResourcesReferencingInformation
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
    if r.RevIncludedProvenanceResourcesReferencingEntity == nil {
        err = errors.New("RevIncluded provenances not requested")
    } else {
        provenances = *r.RevIncludedProvenanceResourcesReferencingEntity
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
    if r.RevIncludedProvenanceResourcesReferencingTarget == nil {
        err = errors.New("RevIncluded provenances not requested")
    } else {
        provenances = *r.RevIncludedProvenanceResourcesReferencingTarget
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
    if r.RevIncludedTaskResourcesReferencingSubject == nil {
        err = errors.New("RevIncluded tasks not requested")
    } else {
        tasks = *r.RevIncludedTaskResourcesReferencingSubject
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
    if r.RevIncludedTaskResourcesReferencingFocus == nil {
        err = errors.New("RevIncluded tasks not requested")
    } else {
        tasks = *r.RevIncludedTaskResourcesReferencingFocus
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
    if r.RevIncludedTaskResourcesReferencingBasedon == nil {
        err = errors.New("RevIncluded tasks not requested")
    } else {
        tasks = *r.RevIncludedTaskResourcesReferencingBasedon
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
    if r.RevIncludedListResourcesReferencingItem == nil {
        err = errors.New("RevIncluded lists not requested")
    } else {
        lists = *r.RevIncludedListResourcesReferencingItem
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
    if r.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
        err = errors.New("RevIncluded evidenceVariables not requested")
    } else {
        evidenceVariables = *r.RevIncludedEvidenceVariableResourcesReferencingSuccessor
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
    if r.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
        err = errors.New("RevIncluded evidenceVariables not requested")
    } else {
        evidenceVariables = *r.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
    if r.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
        err = errors.New("RevIncluded evidenceVariables not requested")
    } else {
        evidenceVariables = *r.RevIncludedEvidenceVariableResourcesReferencingPredecessor
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
    if r.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
        err = errors.New("RevIncluded evidenceVariables not requested")
    } else {
        evidenceVariables = *r.RevIncludedEvidenceVariableResourcesReferencingComposedof
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
    if r.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
        err = errors.New("RevIncluded evidenceVariables not requested")
    } else {
        evidenceVariables = *r.RevIncludedEvidenceVariableResourcesReferencingDependson
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
    if r.RevIncludedObservationResourcesReferencingFocus == nil {
        err = errors.New("RevIncluded observations not requested")
    } else {
        observations = *r.RevIncludedObservationResourcesReferencingFocus
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
    if r.RevIncludedLibraryResourcesReferencingSuccessor == nil {
        err = errors.New("RevIncluded libraries not requested")
    } else {
        libraries = *r.RevIncludedLibraryResourcesReferencingSuccessor
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
    if r.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
        err = errors.New("RevIncluded libraries not requested")
    } else {
        libraries = *r.RevIncludedLibraryResourcesReferencingDerivedfrom
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
    if r.RevIncludedLibraryResourcesReferencingPredecessor == nil {
        err = errors.New("RevIncluded libraries not requested")
    } else {
        libraries = *r.RevIncludedLibraryResourcesReferencingPredecessor
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
    if r.RevIncludedLibraryResourcesReferencingComposedof == nil {
        err = errors.New("RevIncluded libraries not requested")
    } else {
        libraries = *r.RevIncludedLibraryResourcesReferencingComposedof
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
    if r.RevIncludedLibraryResourcesReferencingDependson == nil {
        err = errors.New("RevIncluded libraries not requested")
    } else {
        libraries = *r.RevIncludedLibraryResourcesReferencingDependson
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
    if r.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
        err = errors.New("RevIncluded communicationRequests not requested")
    } else {
        communicationRequests = *r.RevIncludedCommunicationRequestResourcesReferencingBasedon
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
    if r.RevIncludedBasicResourcesReferencingSubject == nil {
        err = errors.New("RevIncluded basics not requested")
    } else {
        basics = *r.RevIncludedBasicResourcesReferencingSubject
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
    if r.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
        err = errors.New("RevIncluded evidences not requested")
    } else {
        evidences = *r.RevIncludedEvidenceResourcesReferencingSuccessor
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
    if r.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
        err = errors.New("RevIncluded evidences not requested")
    } else {
        evidences = *r.RevIncludedEvidenceResourcesReferencingDerivedfrom
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
    if r.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
        err = errors.New("RevIncluded evidences not requested")
    } else {
        evidences = *r.RevIncludedEvidenceResourcesReferencingPredecessor
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
    if r.RevIncludedEvidenceResourcesReferencingComposedof == nil {
        err = errors.New("RevIncluded evidences not requested")
    } else {
        evidences = *r.RevIncludedEvidenceResourcesReferencingComposedof
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
    if r.RevIncludedEvidenceResourcesReferencingDependson == nil {
        err = errors.New("RevIncluded evidences not requested")
    } else {
        evidences = *r.RevIncludedEvidenceResourcesReferencingDependson
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
    if r.RevIncludedAuditEventResourcesReferencingEntity == nil {
        err = errors.New("RevIncluded auditEvents not requested")
    } else {
        auditEvents = *r.RevIncludedAuditEventResourcesReferencingEntity
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
    if r.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
        err = errors.New("RevIncluded conditions not requested")
    } else {
        conditions = *r.RevIncludedConditionResourcesReferencingEvidencedetail
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
    if r.RevIncludedCompositionResourcesReferencingSubject == nil {
        err = errors.New("RevIncluded compositions not requested")
    } else {
        compositions = *r.RevIncludedCompositionResourcesReferencingSubject
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
    if r.RevIncludedCompositionResourcesReferencingEntry == nil {
        err = errors.New("RevIncluded compositions not requested")
    } else {
        compositions = *r.RevIncludedCompositionResourcesReferencingEntry
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
    if r.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
        err = errors.New("RevIncluded detectedIssues not requested")
    } else {
        detectedIssues = *r.RevIncludedDetectedIssueResourcesReferencingImplicated
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
    if r.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
        err = errors.New("RevIncluded questionnaireResponses not requested")
    } else {
        questionnaireResponses = *r.RevIncludedQuestionnaireResponseResourcesReferencingSubject
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
    if r.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
        err = errors.New("RevIncluded clinicalImpressions not requested")
    } else {
        clinicalImpressions = *r.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
    if r.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
        err = errors.New("RevIncluded planDefinitions not requested")
    } else {
        planDefinitions = *r.RevIncludedPlanDefinitionResourcesReferencingSuccessor
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
    if r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
        err = errors.New("RevIncluded planDefinitions not requested")
    } else {
        planDefinitions = *r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
    if r.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
        err = errors.New("RevIncluded planDefinitions not requested")
    } else {
        planDefinitions = *r.RevIncludedPlanDefinitionResourcesReferencingPredecessor
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
    if r.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
        err = errors.New("RevIncluded planDefinitions not requested")
    } else {
        planDefinitions = *r.RevIncludedPlanDefinitionResourcesReferencingComposedof
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
    if r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
        err = errors.New("RevIncluded planDefinitions not requested")
    } else {
        planDefinitions = *r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
    if r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
        err = errors.New("RevIncluded planDefinitions not requested")
    } else {
        planDefinitions = *r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
    }
    return
}

func (r *ResearchDefinitionPlusRelatedResources) GetIncludedResources() map[string]interface{} {
    resourceMap := make(map[string]interface{})
    if r.IncludedLibraryResourcesReferencedByDependsonPath1 != nil {
        for idx := range *r.IncludedLibraryResourcesReferencedByDependsonPath1 {
            rsc := (*r.IncludedLibraryResourcesReferencedByDependsonPath1)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.IncludedLibraryResourcesReferencedByDependsonPath2 != nil {
        for idx := range *r.IncludedLibraryResourcesReferencedByDependsonPath2 {
            rsc := (*r.IncludedLibraryResourcesReferencedByDependsonPath2)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    return resourceMap
}

func (r *ResearchDefinitionPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
    resourceMap := make(map[string]interface{})
    if r.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
        for idx := range *r.RevIncludedAppointmentResourcesReferencingSupportinginfo {
            rsc := (*r.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
        for idx := range *r.RevIncludedEventDefinitionResourcesReferencingSuccessor {
            rsc := (*r.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
        for idx := range *r.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
            rsc := (*r.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
        for idx := range *r.RevIncludedEventDefinitionResourcesReferencingPredecessor {
            rsc := (*r.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
        for idx := range *r.RevIncludedEventDefinitionResourcesReferencingComposedof {
            rsc := (*r.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
        for idx := range *r.RevIncludedEventDefinitionResourcesReferencingDependson {
            rsc := (*r.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedDocumentManifestResourcesReferencingItem != nil {
        for idx := range *r.RevIncludedDocumentManifestResourcesReferencingItem {
            rsc := (*r.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
        for idx := range *r.RevIncludedDocumentManifestResourcesReferencingRelatedref {
            rsc := (*r.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedConsentResourcesReferencingData != nil {
        for idx := range *r.RevIncludedConsentResourcesReferencingData {
            rsc := (*r.RevIncludedConsentResourcesReferencingData)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedMeasureResourcesReferencingSuccessor != nil {
        for idx := range *r.RevIncludedMeasureResourcesReferencingSuccessor {
            rsc := (*r.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
        for idx := range *r.RevIncludedMeasureResourcesReferencingDerivedfrom {
            rsc := (*r.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedMeasureResourcesReferencingPredecessor != nil {
        for idx := range *r.RevIncludedMeasureResourcesReferencingPredecessor {
            rsc := (*r.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedMeasureResourcesReferencingComposedof != nil {
        for idx := range *r.RevIncludedMeasureResourcesReferencingComposedof {
            rsc := (*r.RevIncludedMeasureResourcesReferencingComposedof)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
        for idx := range *r.RevIncludedMeasureResourcesReferencingDependsonPath1 {
            rsc := (*r.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
        for idx := range *r.RevIncludedMeasureResourcesReferencingDependsonPath2 {
            rsc := (*r.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
        for idx := range *r.RevIncludedDocumentReferenceResourcesReferencingRelated {
            rsc := (*r.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
        for idx := range *r.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
            rsc := (*r.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedVerificationResultResourcesReferencingTarget != nil {
        for idx := range *r.RevIncludedVerificationResultResourcesReferencingTarget {
            rsc := (*r.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedContractResourcesReferencingSubject != nil {
        for idx := range *r.RevIncludedContractResourcesReferencingSubject {
            rsc := (*r.RevIncludedContractResourcesReferencingSubject)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
        for idx := range *r.RevIncludedPaymentNoticeResourcesReferencingRequest {
            rsc := (*r.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
        for idx := range *r.RevIncludedPaymentNoticeResourcesReferencingResponse {
            rsc := (*r.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
        for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
            rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
        for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
            rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
        for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
            rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
        for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingComposedof {
            rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
        for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
            rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
        for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
            rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedImplementationGuideResourcesReferencingResource != nil {
        for idx := range *r.RevIncludedImplementationGuideResourcesReferencingResource {
            rsc := (*r.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
        for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
            rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
        for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
            rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
        for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
            rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
        for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
            rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
        for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
            rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
        for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
            rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedCommunicationResourcesReferencingPartof != nil {
        for idx := range *r.RevIncludedCommunicationResourcesReferencingPartof {
            rsc := (*r.RevIncludedCommunicationResourcesReferencingPartof)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedCommunicationResourcesReferencingBasedon != nil {
        for idx := range *r.RevIncludedCommunicationResourcesReferencingBasedon {
            rsc := (*r.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
        for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
            rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
        for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
            rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
        for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
            rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
        for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingComposedof {
            rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
        for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
            rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
        for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
            rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedLinkageResourcesReferencingItem != nil {
        for idx := range *r.RevIncludedLinkageResourcesReferencingItem {
            rsc := (*r.RevIncludedLinkageResourcesReferencingItem)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedLinkageResourcesReferencingSource != nil {
        for idx := range *r.RevIncludedLinkageResourcesReferencingSource {
            rsc := (*r.RevIncludedLinkageResourcesReferencingSource)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
        for idx := range *r.RevIncludedDeviceRequestResourcesReferencingBasedon {
            rsc := (*r.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
        for idx := range *r.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
            rsc := (*r.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
        for idx := range *r.RevIncludedMessageHeaderResourcesReferencingFocus {
            rsc := (*r.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
        for idx := range *r.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
            rsc := (*r.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedProvenanceResourcesReferencingEntity != nil {
        for idx := range *r.RevIncludedProvenanceResourcesReferencingEntity {
            rsc := (*r.RevIncludedProvenanceResourcesReferencingEntity)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedProvenanceResourcesReferencingTarget != nil {
        for idx := range *r.RevIncludedProvenanceResourcesReferencingTarget {
            rsc := (*r.RevIncludedProvenanceResourcesReferencingTarget)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedTaskResourcesReferencingSubject != nil {
        for idx := range *r.RevIncludedTaskResourcesReferencingSubject {
            rsc := (*r.RevIncludedTaskResourcesReferencingSubject)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedTaskResourcesReferencingFocus != nil {
        for idx := range *r.RevIncludedTaskResourcesReferencingFocus {
            rsc := (*r.RevIncludedTaskResourcesReferencingFocus)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedTaskResourcesReferencingBasedon != nil {
        for idx := range *r.RevIncludedTaskResourcesReferencingBasedon {
            rsc := (*r.RevIncludedTaskResourcesReferencingBasedon)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedListResourcesReferencingItem != nil {
        for idx := range *r.RevIncludedListResourcesReferencingItem {
            rsc := (*r.RevIncludedListResourcesReferencingItem)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
        for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
            rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
        for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
            rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
        for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
            rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
        for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingComposedof {
            rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
        for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingDependson {
            rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedObservationResourcesReferencingFocus != nil {
        for idx := range *r.RevIncludedObservationResourcesReferencingFocus {
            rsc := (*r.RevIncludedObservationResourcesReferencingFocus)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedLibraryResourcesReferencingSuccessor != nil {
        for idx := range *r.RevIncludedLibraryResourcesReferencingSuccessor {
            rsc := (*r.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
        for idx := range *r.RevIncludedLibraryResourcesReferencingDerivedfrom {
            rsc := (*r.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedLibraryResourcesReferencingPredecessor != nil {
        for idx := range *r.RevIncludedLibraryResourcesReferencingPredecessor {
            rsc := (*r.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedLibraryResourcesReferencingComposedof != nil {
        for idx := range *r.RevIncludedLibraryResourcesReferencingComposedof {
            rsc := (*r.RevIncludedLibraryResourcesReferencingComposedof)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedLibraryResourcesReferencingDependson != nil {
        for idx := range *r.RevIncludedLibraryResourcesReferencingDependson {
            rsc := (*r.RevIncludedLibraryResourcesReferencingDependson)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
        for idx := range *r.RevIncludedCommunicationRequestResourcesReferencingBasedon {
            rsc := (*r.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedBasicResourcesReferencingSubject != nil {
        for idx := range *r.RevIncludedBasicResourcesReferencingSubject {
            rsc := (*r.RevIncludedBasicResourcesReferencingSubject)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
        for idx := range *r.RevIncludedEvidenceResourcesReferencingSuccessor {
            rsc := (*r.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
        for idx := range *r.RevIncludedEvidenceResourcesReferencingDerivedfrom {
            rsc := (*r.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
        for idx := range *r.RevIncludedEvidenceResourcesReferencingPredecessor {
            rsc := (*r.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedEvidenceResourcesReferencingComposedof != nil {
        for idx := range *r.RevIncludedEvidenceResourcesReferencingComposedof {
            rsc := (*r.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedEvidenceResourcesReferencingDependson != nil {
        for idx := range *r.RevIncludedEvidenceResourcesReferencingDependson {
            rsc := (*r.RevIncludedEvidenceResourcesReferencingDependson)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedAuditEventResourcesReferencingEntity != nil {
        for idx := range *r.RevIncludedAuditEventResourcesReferencingEntity {
            rsc := (*r.RevIncludedAuditEventResourcesReferencingEntity)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
        for idx := range *r.RevIncludedConditionResourcesReferencingEvidencedetail {
            rsc := (*r.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedCompositionResourcesReferencingSubject != nil {
        for idx := range *r.RevIncludedCompositionResourcesReferencingSubject {
            rsc := (*r.RevIncludedCompositionResourcesReferencingSubject)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedCompositionResourcesReferencingEntry != nil {
        for idx := range *r.RevIncludedCompositionResourcesReferencingEntry {
            rsc := (*r.RevIncludedCompositionResourcesReferencingEntry)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
        for idx := range *r.RevIncludedDetectedIssueResourcesReferencingImplicated {
            rsc := (*r.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
        for idx := range *r.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
            rsc := (*r.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
        for idx := range *r.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
            rsc := (*r.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
        for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
            rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
        for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
            rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
        for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
            rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
        for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingComposedof {
            rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
        for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
            rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
        for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
            rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    return resourceMap
}

func (r *ResearchDefinitionPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
    resourceMap := make(map[string]interface{})
    if r.IncludedLibraryResourcesReferencedByDependsonPath1 != nil {
        for idx := range *r.IncludedLibraryResourcesReferencedByDependsonPath1 {
            rsc := (*r.IncludedLibraryResourcesReferencedByDependsonPath1)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.IncludedLibraryResourcesReferencedByDependsonPath2 != nil {
        for idx := range *r.IncludedLibraryResourcesReferencedByDependsonPath2 {
            rsc := (*r.IncludedLibraryResourcesReferencedByDependsonPath2)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
        for idx := range *r.RevIncludedAppointmentResourcesReferencingSupportinginfo {
            rsc := (*r.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
        for idx := range *r.RevIncludedEventDefinitionResourcesReferencingSuccessor {
            rsc := (*r.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
        for idx := range *r.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
            rsc := (*r.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
        for idx := range *r.RevIncludedEventDefinitionResourcesReferencingPredecessor {
            rsc := (*r.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
        for idx := range *r.RevIncludedEventDefinitionResourcesReferencingComposedof {
            rsc := (*r.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
        for idx := range *r.RevIncludedEventDefinitionResourcesReferencingDependson {
            rsc := (*r.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedDocumentManifestResourcesReferencingItem != nil {
        for idx := range *r.RevIncludedDocumentManifestResourcesReferencingItem {
            rsc := (*r.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
        for idx := range *r.RevIncludedDocumentManifestResourcesReferencingRelatedref {
            rsc := (*r.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedConsentResourcesReferencingData != nil {
        for idx := range *r.RevIncludedConsentResourcesReferencingData {
            rsc := (*r.RevIncludedConsentResourcesReferencingData)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedMeasureResourcesReferencingSuccessor != nil {
        for idx := range *r.RevIncludedMeasureResourcesReferencingSuccessor {
            rsc := (*r.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
        for idx := range *r.RevIncludedMeasureResourcesReferencingDerivedfrom {
            rsc := (*r.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedMeasureResourcesReferencingPredecessor != nil {
        for idx := range *r.RevIncludedMeasureResourcesReferencingPredecessor {
            rsc := (*r.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedMeasureResourcesReferencingComposedof != nil {
        for idx := range *r.RevIncludedMeasureResourcesReferencingComposedof {
            rsc := (*r.RevIncludedMeasureResourcesReferencingComposedof)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
        for idx := range *r.RevIncludedMeasureResourcesReferencingDependsonPath1 {
            rsc := (*r.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
        for idx := range *r.RevIncludedMeasureResourcesReferencingDependsonPath2 {
            rsc := (*r.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
        for idx := range *r.RevIncludedDocumentReferenceResourcesReferencingRelated {
            rsc := (*r.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
        for idx := range *r.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
            rsc := (*r.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedVerificationResultResourcesReferencingTarget != nil {
        for idx := range *r.RevIncludedVerificationResultResourcesReferencingTarget {
            rsc := (*r.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedContractResourcesReferencingSubject != nil {
        for idx := range *r.RevIncludedContractResourcesReferencingSubject {
            rsc := (*r.RevIncludedContractResourcesReferencingSubject)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
        for idx := range *r.RevIncludedPaymentNoticeResourcesReferencingRequest {
            rsc := (*r.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
        for idx := range *r.RevIncludedPaymentNoticeResourcesReferencingResponse {
            rsc := (*r.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
        for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
            rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
        for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
            rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
        for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
            rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
        for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingComposedof {
            rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
        for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
            rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
        for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
            rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedImplementationGuideResourcesReferencingResource != nil {
        for idx := range *r.RevIncludedImplementationGuideResourcesReferencingResource {
            rsc := (*r.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
        for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
            rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
        for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
            rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
        for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
            rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
        for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
            rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
        for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
            rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
        for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
            rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedCommunicationResourcesReferencingPartof != nil {
        for idx := range *r.RevIncludedCommunicationResourcesReferencingPartof {
            rsc := (*r.RevIncludedCommunicationResourcesReferencingPartof)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedCommunicationResourcesReferencingBasedon != nil {
        for idx := range *r.RevIncludedCommunicationResourcesReferencingBasedon {
            rsc := (*r.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
        for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
            rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
        for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
            rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
        for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
            rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
        for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingComposedof {
            rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
        for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
            rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
        for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
            rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedLinkageResourcesReferencingItem != nil {
        for idx := range *r.RevIncludedLinkageResourcesReferencingItem {
            rsc := (*r.RevIncludedLinkageResourcesReferencingItem)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedLinkageResourcesReferencingSource != nil {
        for idx := range *r.RevIncludedLinkageResourcesReferencingSource {
            rsc := (*r.RevIncludedLinkageResourcesReferencingSource)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
        for idx := range *r.RevIncludedDeviceRequestResourcesReferencingBasedon {
            rsc := (*r.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
        for idx := range *r.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
            rsc := (*r.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
        for idx := range *r.RevIncludedMessageHeaderResourcesReferencingFocus {
            rsc := (*r.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
        for idx := range *r.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
            rsc := (*r.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedProvenanceResourcesReferencingEntity != nil {
        for idx := range *r.RevIncludedProvenanceResourcesReferencingEntity {
            rsc := (*r.RevIncludedProvenanceResourcesReferencingEntity)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedProvenanceResourcesReferencingTarget != nil {
        for idx := range *r.RevIncludedProvenanceResourcesReferencingTarget {
            rsc := (*r.RevIncludedProvenanceResourcesReferencingTarget)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedTaskResourcesReferencingSubject != nil {
        for idx := range *r.RevIncludedTaskResourcesReferencingSubject {
            rsc := (*r.RevIncludedTaskResourcesReferencingSubject)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedTaskResourcesReferencingFocus != nil {
        for idx := range *r.RevIncludedTaskResourcesReferencingFocus {
            rsc := (*r.RevIncludedTaskResourcesReferencingFocus)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedTaskResourcesReferencingBasedon != nil {
        for idx := range *r.RevIncludedTaskResourcesReferencingBasedon {
            rsc := (*r.RevIncludedTaskResourcesReferencingBasedon)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedListResourcesReferencingItem != nil {
        for idx := range *r.RevIncludedListResourcesReferencingItem {
            rsc := (*r.RevIncludedListResourcesReferencingItem)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
        for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
            rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
        for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
            rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
        for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
            rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
        for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingComposedof {
            rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
        for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingDependson {
            rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedObservationResourcesReferencingFocus != nil {
        for idx := range *r.RevIncludedObservationResourcesReferencingFocus {
            rsc := (*r.RevIncludedObservationResourcesReferencingFocus)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedLibraryResourcesReferencingSuccessor != nil {
        for idx := range *r.RevIncludedLibraryResourcesReferencingSuccessor {
            rsc := (*r.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
        for idx := range *r.RevIncludedLibraryResourcesReferencingDerivedfrom {
            rsc := (*r.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedLibraryResourcesReferencingPredecessor != nil {
        for idx := range *r.RevIncludedLibraryResourcesReferencingPredecessor {
            rsc := (*r.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedLibraryResourcesReferencingComposedof != nil {
        for idx := range *r.RevIncludedLibraryResourcesReferencingComposedof {
            rsc := (*r.RevIncludedLibraryResourcesReferencingComposedof)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedLibraryResourcesReferencingDependson != nil {
        for idx := range *r.RevIncludedLibraryResourcesReferencingDependson {
            rsc := (*r.RevIncludedLibraryResourcesReferencingDependson)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
        for idx := range *r.RevIncludedCommunicationRequestResourcesReferencingBasedon {
            rsc := (*r.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedBasicResourcesReferencingSubject != nil {
        for idx := range *r.RevIncludedBasicResourcesReferencingSubject {
            rsc := (*r.RevIncludedBasicResourcesReferencingSubject)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
        for idx := range *r.RevIncludedEvidenceResourcesReferencingSuccessor {
            rsc := (*r.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
        for idx := range *r.RevIncludedEvidenceResourcesReferencingDerivedfrom {
            rsc := (*r.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
        for idx := range *r.RevIncludedEvidenceResourcesReferencingPredecessor {
            rsc := (*r.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedEvidenceResourcesReferencingComposedof != nil {
        for idx := range *r.RevIncludedEvidenceResourcesReferencingComposedof {
            rsc := (*r.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedEvidenceResourcesReferencingDependson != nil {
        for idx := range *r.RevIncludedEvidenceResourcesReferencingDependson {
            rsc := (*r.RevIncludedEvidenceResourcesReferencingDependson)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedAuditEventResourcesReferencingEntity != nil {
        for idx := range *r.RevIncludedAuditEventResourcesReferencingEntity {
            rsc := (*r.RevIncludedAuditEventResourcesReferencingEntity)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
        for idx := range *r.RevIncludedConditionResourcesReferencingEvidencedetail {
            rsc := (*r.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedCompositionResourcesReferencingSubject != nil {
        for idx := range *r.RevIncludedCompositionResourcesReferencingSubject {
            rsc := (*r.RevIncludedCompositionResourcesReferencingSubject)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedCompositionResourcesReferencingEntry != nil {
        for idx := range *r.RevIncludedCompositionResourcesReferencingEntry {
            rsc := (*r.RevIncludedCompositionResourcesReferencingEntry)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
        for idx := range *r.RevIncludedDetectedIssueResourcesReferencingImplicated {
            rsc := (*r.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
        for idx := range *r.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
            rsc := (*r.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
        for idx := range *r.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
            rsc := (*r.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
        for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
            rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
        for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
            rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
        for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
            rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
        for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingComposedof {
            rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
        for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
            rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    if r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
        for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
            rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
            resourceMap[rsc.Id] = &rsc
        }
    }
    return resourceMap
}
