package mira

import (
	"reflect"
)

// Type is a type info
type Type struct {
	name     string
	v        interface{}
	t        reflect.Type
	pkgPath  string
	nillable bool
	numeric  bool
}

// NewType inspects v and returns a type info
func NewType(v interface{}) *Type {
	t := reflect.TypeOf(v)
	et := t
	for et.Kind() == reflect.Ptr {
		et = et.Elem()
	}

	return &Type{
		name:     et.Name(),
		v:        v,
		t:        et,
		pkgPath:  et.PkgPath(),
		nillable: nillable(t),
		numeric:  numeric(t),
	}
}

// Name is the type name
func (t Type) Name() string {
	return t.name
}

// V is the raw value
func (t Type) V() interface{} {
	return t.v
}

// T is the the reflect.Type
func (t Type) T() reflect.Type {
	return t.t
}

// IsNillable is true when type is nillable and false otherwise
func (t Type) IsNillable() bool {
	return t.nillable
}

// IsNumeric is true when type is numeric
func (t Type) IsNumeric() bool {
	return t.numeric
}

// PkgPath is the package path of the type
func (t Type) PkgPath() string {
	return t.pkgPath
}

func nillable(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Map, reflect.Slice, reflect.Ptr:
		return true
	}
	return false
}

func numeric(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Float32,
		reflect.Float64:
		return true
	case reflect.Ptr:
		et := t.Elem()
		return numeric(et)
	}
	return false
}
