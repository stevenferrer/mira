package mira

import (
	"reflect"
)

// Type is a type info
type Type struct {
	v        interface{}
	t        reflect.Type
	pkgPath  string
	kind     Kind
	nillable bool
}

// NewType inspects v and returns a type info
func NewType(v interface{}) *Type {
	t := reflect.TypeOf(v)
	k := kind(t)
	return &Type{
		v:        v,
		t:        t,
		pkgPath:  pkgPath(t),
		nillable: nillable(k),
		kind:     k,
	}
}

// V returns the value
func (t Type) V() interface{} {
	return t.v
}

// T returns the reflect.Type
func (t Type) T() reflect.Type {
	return t.t
}

// IsNillable is true when type is nillable and false otherwise
func (t Type) IsNillable() bool {
	return t.nillable
}

func (t Type) Kind() Kind {
	return t.kind
}

// PkgPath is the package path of the type
func (t Type) PkgPath() string {
	return t.pkgPath
}

func nillable(k Kind) bool {
	switch k {
	case Slice, Map, Ptr:
		return true
	}
	return false
}

func kind(t reflect.Type) Kind {
	if numeric(t) {
		return Numeric
	}

	switch t.Kind() {
	case reflect.Bool:
		return Bool
	case reflect.String:
		return String
	case reflect.Slice:
		return Slice
	case reflect.Array:
		return Array
	case reflect.Map:
		return Map
	case reflect.Struct:
		return Struct
	case reflect.Ptr:
		return Ptr
	}

	return Unknown
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
	}

	return false
}

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
