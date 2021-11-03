package mira_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stevenferrer/mira"
)

type ByteArray [16]byte

func TestNewTypeInfo(t *testing.T) {
	type args struct{ v interface{} }
	type want struct {
		name,
		pkgPath string
		nillable,
		numeric bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "int",
			args: args{
				v: int(0),
			},
			want: want{
				name:    "int",
				numeric: true,
			},
		},
		{
			name: "ptr to int",
			args: args{
				v: mira.IntPtr(0),
			},
			want: want{
				nillable: true,
				name:     "int",
			},
		},
		{
			name: "bool",
			args: args{
				v: false,
			},
			want: want{
				name: "bool",
			},
		},
		{
			name: "bool ptr",
			args: args{
				v: mira.BoolPtr(false),
			},
			want: want{
				nillable: true,
				name:     "bool",
			},
		},
		{
			name: "string",
			args: args{
				v: "",
			},
			want: want{
				name: "string",
			},
		},
		{
			name: "string ptr",
			args: args{
				v: mira.StrPtr(""),
			},
			want: want{
				nillable: true,
				name:     "string",
			},
		},
		{
			name: "map",
			args: args{
				v: map[string]interface{}{},
			},
			want: want{
				nillable: true,
			},
		},
		{
			name: "map ptr",
			args: args{
				v: &map[string]interface{}{},
			},
			want: want{
				nillable: true,
			},
		},
		{
			name: "big.Int",
			args: args{
				v: big.Int{},
			},
			want: want{
				pkgPath: "math/big",
				name:    "Int",
			},
		},
		{
			name: "big.Int ptr",
			args: args{
				v: &big.Int{},
			},
			want: want{
				pkgPath:  "math/big",
				nillable: true,
				name:     "Int",
			},
		},

		{
			name: "int slice",
			args: args{
				v: []int{},
			},
			want: want{
				nillable: true,
				name:     "int",
			},
		},
		{
			name: "int array",
			args: args{
				v: [10]int{},
			},
			want: want{
				name: "int",
			},
		},
		{
			name: "big.Int slice",
			args: args{
				v: []big.Int{},
			},
			want: want{
				nillable: true,
				pkgPath:  "math/big",
				name:     "Int",
			},
		},
		{
			name: "big.Int ptr slice",
			args: args{
				v: []*big.Int{},
			},
			want: want{
				nillable: true,
				pkgPath:  "math/big",
				name:     "Int",
			},
		},
		{
			name: "big.Int array",
			args: args{
				v: [10]big.Int{},
			},
			want: want{
				pkgPath: "math/big",
				name:    "Int",
			},
		},
		{
			name: "big.Int ptr array",
			args: args{
				v: [10]*big.Int{},
			},
			want: want{
				pkgPath: "math/big",
				name:    "Int",
			},
		},
		{
			name: "ByteArray",
			args: args{
				v: ByteArray{},
			},
			want: want{
				pkgPath: "github.com/stevenferrer/mira_test",
				name:    "ByteArray",
			},
		},
		{
			name: "ByteArray ptr",
			args: args{
				v: &ByteArray{},
			},
			want: want{
				pkgPath:  "github.com/stevenferrer/mira_test",
				nillable: true,
				name:     "ByteArray",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := mira.NewTypeInfo(tc.args.v)
			assert.Equal(t, tc.want.numeric, got.IsNumeric(), "expected numeric")
			assert.Equal(t, tc.want.nillable, got.IsNillable(), "expected nillable")
			assert.Equal(t, tc.want.pkgPath, got.PkgPath(), "expected pkgPath")
			assert.Equal(t, tc.want.name, got.Name(), "expected type name")
			assert.NotNil(t, got.V())
			assert.NotNil(t, got.T())
		})
	}
}
