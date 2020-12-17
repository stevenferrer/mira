package mira

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewType(t *testing.T) {
	type args struct {
		v interface{}
	}

	bPtr := BoolPtr(false)
	sPtr := StrPtr("")
	ints := []int{}
	inta := [16]int{}
	bigIntPtr := big.NewInt(0)
	bigInts := []big.Int{}
	bigIntPtrs := []*big.Int{}
	bigInta := [10]big.Int{}
	bigIntPtra := [10]*big.Int{}
	m := map[string]interface{}{}

	myUUID := uuid.New()

	tests := []struct {
		name string
		args args
		want *Type
	}{
		{
			name: "int",
			args: args{
				v: int(0),
			},
			want: &Type{
				v:    int(0),
				kind: Numeric,
				name: "int",
			},
		},
		{
			name: "*int",
			args: args{
				v: IntPtr(0),
			},
			want: &Type{
				v:        IntPtr(0),
				nillable: true,
				kind:     Ptr,
				name:     "int",
			},
		},
		{
			name: "bool",
			args: args{
				v: false,
			},
			want: &Type{
				v:    false,
				kind: Bool,
				name: "bool",
			},
		},
		{
			name: "*bool",
			args: args{
				v: bPtr,
			},
			want: &Type{
				v:        bPtr,
				kind:     Ptr,
				nillable: true,
				name:     "bool",
			},
		},
		{
			name: "string",
			args: args{
				v: *sPtr,
			},
			want: &Type{
				v:    *sPtr,
				kind: String,
				name: "string",
			},
		},
		{
			name: "*string",
			args: args{
				v: sPtr,
			},
			want: &Type{
				v:        sPtr,
				kind:     Ptr,
				nillable: true,
				name:     "string",
			},
		},
		{
			name: "map",
			args: args{
				v: m,
			},
			want: &Type{
				v:        m,
				kind:     Map,
				nillable: true,
			},
		},
		{
			name: "*map",
			args: args{
				v: &m,
			},
			want: &Type{
				v:        &m,
				kind:     Ptr,
				nillable: true,
			},
		},
		{
			name: "big.Int",
			args: args{
				v: *bigIntPtr,
			},
			want: &Type{
				v:       *bigIntPtr,
				pkgPath: "math/big",
				kind:    Struct,
				name:    "Int",
			},
		},
		{
			name: "*big.Int",
			args: args{
				v: bigIntPtr,
			},
			want: &Type{
				v:        bigIntPtr,
				pkgPath:  "math/big",
				kind:     Ptr,
				nillable: true,
				name:     "Int",
			},
		},

		{
			name: "[]int",
			args: args{
				v: ints,
			},
			want: &Type{
				v:        ints,
				nillable: true,
				kind:     Slice,
				name:     "int",
			},
		},
		{
			name: "[16]int",
			args: args{
				v: inta,
			},
			want: &Type{
				v:    inta,
				kind: Array,
				name: "int",
			},
		},
		{
			name: "[]big.Int",
			args: args{
				v: bigInts,
			},
			want: &Type{
				v:        bigInts,
				nillable: true,
				pkgPath:  "math/big",
				kind:     Slice,
				name:     "Int",
			},
		},
		{
			name: "[]*big.Int",
			args: args{
				v: bigIntPtrs,
			},
			want: &Type{
				v:        bigIntPtrs,
				nillable: true,
				pkgPath:  "math/big",
				kind:     Slice,
				name:     "Int",
			},
		},
		{
			name: "[10]big.Int",
			args: args{
				v: bigInta,
			},
			want: &Type{
				v:       bigInta,
				pkgPath: "math/big",
				kind:    Array,
				name:    "Int",
			},
		},
		{
			name: "[]*big.Int",
			args: args{
				v: bigIntPtra,
			},
			want: &Type{
				v:       bigIntPtra,
				pkgPath: "math/big",
				kind:    Array,
				name:    "Int",
			},
		},
		{
			name: "uuid.UUID",
			args: args{
				v: myUUID,
			},
			want: &Type{
				v:       myUUID,
				pkgPath: "github.com/google/uuid",
				kind:    Array,
				name:    "UUID",
			},
		},
		{
			name: "*uuid.UUID",
			args: args{
				v: &myUUID,
			},
			want: &Type{
				v:        &myUUID,
				pkgPath:  "github.com/google/uuid",
				kind:     Ptr,
				nillable: true,
				name:     "UUID",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewType(tt.args.v)
			assert.Equal(t, tt.want.V(), got.V(), "expected value")
			assert.Equal(t, tt.want.IsNillable(), got.IsNillable(), "expected nillable")
			assert.Equal(t, tt.want.Kind().String(), got.Kind().String(), "expected kind")
			assert.Equal(t, tt.want.PkgPath(), got.PkgPath(), "expected pkg path")
			assert.Equal(t, tt.want.Name(), got.Name(), "expected type name")
			assert.NotNil(t, got.t)
		})
	}
}

func Test_name(t *testing.T) {
	type args struct {
		t reflect.Type
	}

	i := int(0)
	s := ""
	ints := []int{}
	inta := [8]int{}
	intsPtr := []*int{}
	intaPtr := [8]*int{}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "int",
			args: args{
				t: reflect.TypeOf(i),
			},
			want: "int",
		},
		{
			name: "*int",
			args: args{
				t: reflect.TypeOf(&i),
			},
			want: "int",
		},
		{
			name: "string",
			args: args{
				t: reflect.TypeOf(s),
			},
			want: "string",
		},
		{
			name: "*string",
			args: args{
				t: reflect.TypeOf(&s),
			},
			want: "string",
		},
		{
			name: "[]int",
			args: args{
				t: reflect.TypeOf(ints),
			},
			want: "int",
		},
		{
			name: "[8]int",
			args: args{
				t: reflect.TypeOf(inta),
			},
			want: "int",
		},
		{
			name: "[]*int",
			args: args{
				t: reflect.TypeOf(intsPtr),
			},
			want: "int",
		},
		{
			name: "[8]*int",
			args: args{
				t: reflect.TypeOf(intaPtr),
			},
			want: "int",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s name is %s", tt.name, tt.want), func(t *testing.T) {
			got := name(tt.args.t)
			assert.Equal(t, tt.want, got)
		})
	}
}
