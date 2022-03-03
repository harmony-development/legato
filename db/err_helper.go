package db

import "fmt"

func Wrap(e error, msg string) error {
	return fmt.Errorf("%s: %w", msg, e)
}

func TryWrap(e error, msg string) error {
	if e == nil {
		return nil
	}

	return Wrap(e, msg)
}
