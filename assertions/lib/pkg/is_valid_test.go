package assertions

import (
	"reflect"
	"testing"
)

func Test_findTypeFrom(t *testing.T) {
	type MyInterface interface{}

	tests := []struct {
		name string
		when func() reflect.Type
		want reflect.Type
	}{
		{
			name: "MyInterface should return the type of a pointer to MyInterface",
			when: func() reflect.Type {
				return findTypeFrom[MyInterface]()
			},
			want: reflect.TypeOf(new(MyInterface)),
		},
		{
			name: "string should return the type of string",
			when: func() reflect.Type {
				return findTypeFrom[string]()
			},
			want: reflect.TypeOf(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.when(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findTypeFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
