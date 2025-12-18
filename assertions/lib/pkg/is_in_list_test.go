package assertions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type tint int

func (i tint) String() string { return fmt.Sprintf("{%d}", int(i)) }

func TestIsNotIn(t *testing.T) {
	type args[T comparable] struct {
		v    T
		list []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want assert.ErrorAssertionFunc
	}
	tests := []testCase[string]{
		{
			name: "If nil list is provided, then returns false",
			args: args[string]{v: "a", list: nil},
			want: assert.Error,
		},
		{
			name: "If a is in the list, then returns true",
			args: args[string]{v: "a", list: []string{"a"}},
			want: assert.NoError,
		},
		{
			name: "If a is NOT in the list, then returns false",
			args: args[string]{v: "a", list: []string{"b"}},
			want: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsIn(tt.args.v, tt.args.list...)
			tt.want(t, got)
		})
	}
}

func TestIsNotIn2(t *testing.T) {
	type args[T comparable] struct {
		v    T
		list []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want assert.ErrorAssertionFunc
	}
	tests := []testCase[tint]{
		{
			name: "If nil list is provided, then returns false",
			args: args[tint]{v: 1, list: nil},
			want: func(t assert.TestingT, err error, args ...interface{}) bool {
				return assert.Error(t, err) &&
					assert.Contains(t, "value=[{1}] not in the permitted value list=[]", err.Error())
			},
		},
		{
			name: "If 1 is in the list, then returns true",
			args: args[tint]{v: 1, list: []tint{1}},
			want: assert.NoError,
		},
		{
			name: "If 1 is NOT in the list, then returns false",
			args: args[tint]{v: 1, list: []tint{2}},
			want: func(t assert.TestingT, err error, args ...interface{}) bool {
				return assert.Error(t, err) &&
					assert.Contains(t, "value=[{1}] not in the permitted value list=[{2}]", err.Error())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsIn(tt.args.v, tt.args.list...)
			tt.want(t, got)
		})
	}
}
