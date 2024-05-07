package errorswrapper

import (
	"errors"
	"fmt"
)

func Wrap(err error, message string) error {
	if err == nil {
		return errors.New(message)
	}
	newErr := fmt.Errorf("%s -> %w", message, err)
	return newErr
}
