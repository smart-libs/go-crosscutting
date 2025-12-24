package playground

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	assertions "github.com/smart-libs/go-crosscutting/assertions/lib/pkg"
	"github.com/smart-libs/go-crosscutting/serror/lib/pkg"
)

func wrapError(objectNameToBeValidated string, val any, err error) error {
	switch castedErr := err.(type) {
	case nil:
		return nil
	case *validator.InvalidValidationError:
		return serror.WrapAsInternalError(err)
	case validator.ValidationErrors:
		var causes []error
		for _, innerErr := range castedErr {
			if innerErr.Namespace() != "" {
				causes = append(causes, errors.New(innerErr.Namespace()))
			} else if innerErr.Field() != "" {
				causes = append(causes, errors.New(innerErr.Field()))
			} else {
				causes = append(causes, fmt.Errorf("validation: %s", innerErr.Tag()))
			}
		}
		return assertions.WrapAsIllegalArgumentValueWithCause(objectNameToBeValidated, val, causes...)
	default:
		return assertions.WrapError(objectNameToBeValidated, val, err)
	}
}
