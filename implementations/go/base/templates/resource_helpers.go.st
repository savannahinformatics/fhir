package models

import "reflect"

// NewStructForResourceName initializes a struct pointer with a specific name
func NewStructForResourceName(name string) interface{} {
	rStruct := StructForResourceName(name)
	structPtr := reflect.New(reflect.TypeOf(rStruct))
	structPtr.Elem().Set(reflect.ValueOf(rStruct))
	return structPtr.Interface()
}

// SliceForResourceName initializes a slice with the given resource's name
func SliceForResourceName(name string, len int, cap int) interface{} {
	rType := reflect.TypeOf(StructForResourceName(name))
	return reflect.MakeSlice(reflect.SliceOf(rType), len, cap).Interface()
}

// NewSliceForResourceName creates a pointer to a slice pointer
func NewSliceForResourceName(name string, len int, cap int) interface{} {
	rSlice := SliceForResourceName(name, len, cap)
	rSlicePtr := reflect.New(reflect.TypeOf(rSlice))
	rSlicePtr.Elem().Set(reflect.ValueOf(rSlice))
	return rSlicePtr.Interface()
}

// StructForResourceName creates a struct given a resource name
func StructForResourceName(name string) interface{} {
	switch name {
	<Resources:{r |	case "<r>":
	return <r><\u007B><\u007D>
	}>
	default:
		return nil
	}
}

// SlicePlusForResourceName initializes a sized slice for a resource name
func SlicePlusForResourceName(name string, len int, cap int) interface{} {
	rType := reflect.TypeOf(StructPlusForResourceName(name))
	return reflect.MakeSlice(reflect.SliceOf(rType), len, cap).Interface()
}

// NewSlicePlusForResourceName initializes a sized slice of pointers
func NewSlicePlusForResourceName(name string, len int, cap int) interface{} {
	rSlice := SlicePlusForResourceName(name, len, cap)
	rSlicePtr := reflect.New(reflect.TypeOf(rSlice))
	rSlicePtr.Elem().Set(reflect.ValueOf(rSlice))
	return rSlicePtr.Interface()
}

// StructPlusForResourceName generates a struct for a resource name
func StructPlusForResourceName(name string) interface{} {
	switch name {
	<Resources:{r |	case "<r>":
	return <r>Plus<\u007B><\u007D>
	}>
	default:
		return nil
	}
}

// PluralizeLowerResourceName generates plural resource names
func PluralizeLowerResourceName(name string) string {
	switch name {
	<Resources:{r |	case "<r>":
	return "<Plurals.(r)>"
	}>
	default:
		return ""
	}
}
