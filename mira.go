package mira

import (
	"reflect"
)

// Type is a simplified type info
type Type struct {
	name     string
	v        interface{}
	t        reflect.Type
	pkgPath  string
	nillable bool
}

// NewType inspects v and gives *Type
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
	}
}

// Name is the type name
func (t Type) Name() string {
	return t.name
}

// V is the type value
func (t Type) V() interface{} {
	return t.v
}

// T is the the reflect.Type
func (t Type) T() reflect.Type {
	return t.t
}

// Nillable is true when type is nillable and false otherwise
func (t Type) Nillable() bool {
	return t.nillable
}

// PkgPath is the package path of the type
func (t Type) PkgPath() string {
	return t.pkgPath
}

func nillable(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Slice, reflect.Array,
		reflect.Ptr, reflect.Map:
		return true
	}
	return false
}
