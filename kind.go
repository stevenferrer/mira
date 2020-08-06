package mira

import "strconv"

// Kind is a simplified kind
type Kind uint

// List of type kinds
const (
	// Unknown is an unknow kind
	Unknown Kind = iota
	// Bool is a boolean kind
	Bool
	// Numeric is a numeric kind which includes int(8, 16, 32, 64), uint(8, 16, 32, 64), float(32, 64)
	Numeric
	// Array is an array kind
	Array
	// Chan is a channel kind
	Chan
	// Func is a func kind
	Func
	// Interface is an interface kind
	Interface
	// Map is a map kind
	Map
	// Ptr is a pointer kind
	Ptr
	// Slice is a slice kind
	Slice
	// String is a string kind
	String
	// Struct is a struct kind
	Struct
)

// String returns the name of k.
func (k Kind) String() string {
	if int(k) < len(kindNames) {
		return kindNames[k]
	}
	return "kind" + strconv.Itoa(int(k))
}

var kindNames = []string{
	Unknown:   "invalid",
	Numeric:   "numeric",
	Bool:      "bool",
	Array:     "array",
	Chan:      "chan",
	Func:      "func",
	Interface: "interface",
	Map:       "map",
	Ptr:       "ptr",
	Slice:     "slice",
	String:    "string",
	Struct:    "struct",
}
