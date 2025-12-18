package playground

import (
	"errors"
	"github.com/joomcode/errorx"
	serror "github.com/smart-libs/go-crosscutting/serror/lib/pkg"
	"github.com/stretchr/testify/assert"
	assertions "githup.com/smart-libs/go-crosscutting/assertions/lib/pkg"
	"testing"
)

func TestIsValid(t *testing.T) {
	type Email string
	assertions.RegisterPerType[Email](func(email Email) error {
		return IsValidWithTag(string(email), "email", "required,email")
	})

	type args struct {
		objectName string
	}
	type testCase struct {
		name string
		args args
		when func(args) error
		want func(t *testing.T, r error)
	}
	tests := []testCase{
		{
			name: "I can validatealidate email",
			args: args{
				objectName: "my-name",
			},
			when: func(args args) error {
				return assertions.IsValid(Email("test@route.com"), args.objectName)
			},
			want: func(t *testing.T, r error) {
				assert.Equal(t, nil, r)
			},
		},
		{
			name: "I receive an error if it is not an email",
			args: args{
				objectName: "my-name",
			},
			when: func(args args) error {
				return assertions.IsValid[Email]("route.com", args.objectName)
			},
			want: func(t *testing.T, r error) {
				if assert.NotEqual(t, nil, r) {
					assert.Truef(t, errors.Is(r, assertions.IllegalArgument{}), "error should be IllegalArgument")
					assert.Equal(t, "illegal_argument: invalid value=[route.com] for argument=[my-name], caused by validation: email",
						r.Error())
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.when(tt.args)
			tt.want(t, got)
		})
	}
}

func TestIsValidWithTag(t *testing.T) {
	type args struct {
		objectName string
		tag        string
	}
	type testCase struct {
		name string
		args args
		when func(args) error
		want func(t *testing.T, r error)
	}

	assertions.WrapAsIllegalArgumentValueWithCause = serror.IllegalArgumentValueWithCause

	tests := []testCase{
		{
			name: "I can mutableValidate email",
			args: args{
				objectName: "my-name",
				tag:        "email",
			},
			when: func(args args) error {
				return IsValidWithTag("test@route.com", args.objectName, args.tag)
			},
			want: func(t *testing.T, r error) {
				assert.Equal(t, nil, r)
			},
		},
		{
			name: "I receive an error if it is not an email",
			args: args{
				objectName: "my-name",
				tag:        "email",
			},
			when: func(args args) error {
				return IsValidWithTag("route.com", args.objectName, args.tag)
			},
			want: func(t *testing.T, r error) {
				if assert.NotEqual(t, nil, r) {
					assert.Truef(t, errorx.IsOfType(r, errorx.IllegalArgument), "error should be errorx.IllegalArgument")
					assert.Equal(t, "common.illegal_argument: Illegal value=[my-name] for argument=[route.com], cause: validation: email",
						r.Error())
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.when(tt.args)
			tt.want(t, got)
		})
	}
}

func TestRegisterPerType(t *testing.T) {
	type EMail string

	validFunc := func(e EMail) error {
		return mutableValidate.Var(string(e), "required,email")
	}

	emailValidator := assertions.NewValidatorWithFunc("email-type", validFunc)
	assert.Error(t, emailValidator.Validate("marcelo@route"))

	assertions.RegisterPerType(validFunc)
	RegisterPerTag("email", validFunc)

	assert.NoError(t, mutableValidate.Var(EMail("marcelo@route.com"), "email"))
	assert.Error(t, mutableValidate.Var(EMail("marcelo@route"), "email"))
	assert.Error(t, assertions.IsValid[EMail]("marcelo@route", "email"))
	assert.NoError(t, assertions.IsValid(EMail("marcelo@route.com"), "email"))
}
