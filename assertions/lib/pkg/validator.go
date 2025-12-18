package assertions

type (
	Validator[T any] interface {
		Validate(T) error
		// MustValidate invokes Validate and panic if it returns error
		MustValidate(T)
	}

	defaultValidator[T any] struct {
		validateFunc     func(T) error
		targetObjectName string
	}
)

func NewCompositeValidator[T any](validators ...Validator[T]) Validator[T] {
	return nil
}

func NewValidatorWithFunc[T any](objectNameToBeValidated string, validate func(T) error) Validator[T] {
	return defaultValidator[T]{
		validateFunc:     validate,
		targetObjectName: objectNameToBeValidated,
	}
}

func NewValidatorWithTag[T any](objectNameToBeValidated T, tag string) Validator[T] {
	return nil
}

func (d defaultValidator[T]) Validate(v T) error {
	return WrapError(d.targetObjectName, v, d.validateFunc(v))
}

func (d defaultValidator[T]) MustValidate(v T) {
	Must(d.Validate(v))
}
