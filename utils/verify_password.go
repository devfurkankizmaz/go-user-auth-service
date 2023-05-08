package utils

import (
	"errors"
)

var err = errors.New("passwords did not match")

func VerifyPassword(password string, confirmpass string) error {
	if password != confirmpass {
		return err
	}
	return nil
}
