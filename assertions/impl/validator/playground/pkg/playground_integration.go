package playground

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	assertions "github.com/smart-libs/go-crosscutting/assertions/lib/pkg"
	"reflect"
)

// This source provides an integration between go-assertions and github.com/go-playground/validator/v10
// that should not be used directly.

type (
	IntegratorValidator struct {
		validatorFuncPerType assertions.ValidatorFuncPerType
		tag                  string
	}
)

var (
	// readOnlyValidate has an instance of validator.Validate that is never modified which works
	// as a secure place to invoke original validations
	readOnlyValidate *validator.Validate

	// mutableValidate is used to hold all the new tags created through go-assertions that will
	// invoke the readOnlyValidate validation to avoid any tag overwriting.
	mutableValidate *validator.Validate

	validationFuncPerTag = map[string]IntegratorValidator{}
)

func (i IntegratorValidator) validate(fl validator.FieldLevel) bool {
	val := fl.Field().Interface()
	valType := fl.Field().Type()
	if funcFound := i.validatorFuncPerType.GetForType(valType); funcFound != nil {
		return funcFound(val) == nil
	}
	return readOnlyValidate.Var(val, i.tag) == nil
}

func init() {
	mutableValidate = validator.New(validator.WithRequiredStructEnabled())
	readOnlyValidate = validator.New(validator.WithRequiredStructEnabled())

	assertions.RegisterPerKind(reflect.Struct, isValidStruct)
}

func RegisterPerTag[T any](tag string, newFunc assertions.ValidationFunc[T]) assertions.ValidationFunc[T] {
	integrator, found := validationFuncPerTag[tag]
	if !found {
		integrator = IntegratorValidator{
			tag:                  tag,
			validatorFuncPerType: assertions.ValidatorFuncPerType{},
		}
		validationFuncPerTag[tag] = integrator
		err := mutableValidate.RegisterValidation(tag, integrator.validate, true)
		if err != nil {
			panic(err)
		}
	}

	return assertions.RegisterPerTypeWith[T](integrator.validatorFuncPerType, func(v T) error {
		if err := newFunc(v); err != nil {
			return assertions.WrapAsInternalError(fmt.Errorf("tag=[%s] %w", tag, err))
		}
		return nil
	})
}

func isValidStruct(val any) error {
	return mutableValidate.Struct(val)
}

// IsValidWithTag is a function that uses the validation function associated with the given tag to
// validate the given value. It does not use the registry used by IsValid.
func IsValidWithTag[T any](val T, objectName, tag string) error {
	return wrapError(objectName, val, mutableValidate.Var(val, tag))
}
