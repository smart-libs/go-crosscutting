package serror

import (
	"fmt"
	"github.com/joomcode/errorx"
)

var (
	// Start: Traits used to identify the root cause domain

	// UserRootCauseTrait to inform it is an error cause by user
	UsrRootCauseTrait = errorx.RegisterTrait("traitUsrRootCause")
	// AppRootCauseTrait to inform it is an error cause by user
	AppRootCauseTrait = errorx.RegisterTrait("traitAppRootCause")
	// CmpRootCauseTrait to inform it is an error cause by a component, when the software is intent to be a component
	CmpRootCauseTrait = errorx.RegisterTrait("traitCmpRootCause")
	// ExtRootCauseTrait to inform the error was caused by an external dependency like component, hardware, other app
	ExtRootCauseTrait = errorx.RegisterTrait("traitExtRootCause")
	// End: Traits used to identify the root cause domain

	// Errors to wrap traits

	UsrError = errorx.CommonErrors.NewType("user_error", UsrRootCauseTrait).
			ApplyModifiers(errorx.TypeModifierOmitStackTrace)
	AppError = errorx.CommonErrors.NewType("app_error", AppRootCauseTrait).
			ApplyModifiers(errorx.TypeModifierOmitStackTrace)
	CmpError = errorx.CommonErrors.NewType("component_error", CmpRootCauseTrait).
			ApplyModifiers(errorx.TypeModifierOmitStackTrace)
	ExtError = errorx.CommonErrors.NewType("external_error", ExtRootCauseTrait).
			ApplyModifiers(errorx.TypeModifierOmitStackTrace)

	// IllegalConfig is a type for invalid argument error
	IllegalConfig = errorx.CommonErrors.NewType("illegal_config")

	NotFoundError  = errorx.CommonErrors.NewType("not_found_error", errorx.NotFound())
	DuplicateError = errorx.CommonErrors.NewType("duplicate_error", errorx.Duplicate())

	Temporary = errorx.CommonErrors.NewType("temporary_error", errorx.Temporary())
)

func WrapAsUsrError(err error) error {
	return UsrError.WrapWithNoMessage(err)
}

func WrapAsAppError(err error) error {
	return AppError.WrapWithNoMessage(err)
}

func WrapAsCmpError(err error) error {
	return CmpError.WrapWithNoMessage(err)
}

func WrapAsExtError(err error) error {
	return ExtError.WrapWithNoMessage(err)
}

func IllegalArgumentValue[T any](paramName string, paramValue T) error {
	return errorx.IllegalArgument.New("Illegal value=[%v] for argument=[%v]", paramName, paramValue)
}

func IllegalArgumentValueWithCause[T any](paramName string, paramValue T, causes ...error) error {
	if len(causes) == 0 {
		return IllegalArgumentValue(paramName, paramValue)
	}
	return errorx.WrapMany(errorx.IllegalArgument,
		fmt.Sprintf("Illegal value=[%v] for argument=[%v]", paramName, paramValue),
		causes...)
}

func IllegalConfigParamValue[T any](paramName string, paramValue T) error {
	return IllegalConfig.New("Illegal value=[%v] for config parameter=[%v]", paramName, paramValue)
}

func WrapAsTemporary(err error) error {
	if errorx.IsTemporary(err) {
		return err
	}

	return Temporary.WrapWithNoMessage(err)
}

func WrapAsTimeout(err error) error {
	if errorx.IsTimeout(err) {
		return err
	}

	return errorx.TimeoutElapsed.WrapWithNoMessage(err)
}

func WrapAsInternalError(err error) error {
	if errorx.IsOfType(err, errorx.InternalError) {
		return err
	}

	return errorx.InternalError.WrapWithNoMessage(err)
}
