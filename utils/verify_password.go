package utils

import (
	"errors"
)

func VerifyPassword(password string, confirmpass string) error {
	err := errors.New("passwords did not match")
	if password != confirmpass {
		return err
	}
	return nil
}
