package mira

import (
	"math/big"
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
			assert.NotNil(t, got.t)
		})
	}
}
