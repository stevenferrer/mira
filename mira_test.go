package mira

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewType(t *testing.T) {
	type args struct {
		v interface{}
	}

	mt := &Type{}
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
				name:     "int",
				v:        int(0),
				t:        reflect.TypeOf(int(0)),
				nillable: false,
				numeric:  true,
				pkgPath:  "",
			},
		},
		{
			name: "*int",
			args: args{
				v: IntPtr(0),
			},
			want: &Type{
				name:     "int",
				v:        IntPtr(0),
				t:        reflect.TypeOf(IntPtr(0)).Elem(),
				nillable: true,
				numeric:  true,
				pkgPath:  "",
			},
		},
		{
			name: "mira.Type",
			args: args{
				v: Type{},
			},
			want: &Type{
				name:     "Type",
				v:        Type{},
				t:        reflect.TypeOf(Type{}),
				nillable: false,
				pkgPath:  "github.com/sf9v/mira",
			},
		},
		{
			name: "*mira.Type",
			args: args{
				v: mt,
			},
			want: &Type{
				name:     "Type",
				v:        mt,
				t:        reflect.TypeOf(mt).Elem(),
				nillable: true,
				pkgPath:  "github.com/sf9v/mira",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewType(tt.args.v)
			assert.Equal(t, tt.want.Name(), got.Name())
			assert.Equal(t, tt.want.V(), got.V())
			assert.Equal(t, tt.want.T(), got.T())
			assert.Equal(t, tt.want.IsNillable(), got.IsNillable())
			assert.Equal(t, tt.want.IsNumeric(), got.IsNumeric())
			assert.Equal(t, tt.want.PkgPath(), got.PkgPath())
		})
	}
}
