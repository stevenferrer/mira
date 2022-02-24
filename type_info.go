package mira

import (
	"reflect"
)

// TypeInfo is a type info
type TypeInfo struct {
	v interface{}
	t reflect.Type
}

// NewTypeInfo instrospects v and returns a TypeInfo
func NewTypeInfo(v interface{}) TypeInfo {
	return TypeInfo{v: v, t: reflect.TypeOf(v)}
}

// V returns the interface value
func (ti TypeInfo) V() interface{} {
	return ti.v
}

// T returns the reflect.Type
func (ti TypeInfo) T() reflect.Type {
	return ti.t
}

// Name returns the name of the type within
// its package for a defined type.
//
// This also works with arrays, slices and pointers.
func (ti TypeInfo) Name() string {
	return name(ti.t)
}

// IsNillable returns true if the type is
// nillable (i.e. slices, maps and pointers)
func (ti TypeInfo) IsNillable() bool {
	return isNillable(ti.t)
}

// IsNumeric returns true if the type is
// numeric i.e. int, uint or float.
func (ti TypeInfo) IsNumeric() bool {
	return isNumeric(ti.t)
}

// PkgPath returns a defined type's package path.
//
// This also works with arrays, slices, maps and pointers.
func (ti TypeInfo) PkgPath() string {
	return pkgPath(ti.t)
}

// isNillable returns true when t is nillable.
func isNillable(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Slice, reflect.Map, reflect.Ptr:
		return true
	}
	return false
}

// isNumeric returns true when t is a numeric type.
func isNumeric(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Int, reflect.Int8,
		reflect.Int16, reflect.Int32,
		reflect.Int64, reflect.Uint,
		reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return true
	}

	return false
}

// pkgPath recursively returns a defined type's package path.
func pkgPath(t reflect.Type) string {
	pkg := t.PkgPath()
	if len(pkg) > 0 {
		return pkg
	}

	switch t.Kind() {
	case reflect.Ptr, reflect.Map,
		reflect.Slice, reflect.Array:
		return pkgPath(t.Elem())
	}
	return pkg
}

// name recursively returns the type's name within its
// package for a defined type.
func name(t reflect.Type) string {
	tn := t.Name()
	if tn != "" {
		return tn
	}

	switch t.Kind() {
	case reflect.Slice, reflect.Array, reflect.Ptr:
		return name(t.Elem())
	}

	return tn
}
