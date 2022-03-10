package db

import "fmt"

func Wrap(e error, msg string, args ...interface{}) error {
	return fmt.Errorf("%s: %w", fmt.Sprintf(msg, args...), e)
}

func TryWrap(e error, msg string, args ...interface{}) error {
	if e == nil {
		return nil
	}

	return Wrap(e, msg, args...)
}
