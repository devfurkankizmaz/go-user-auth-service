package utils

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

func EmailValid(mail string) error {
	err := errors.New("wrong email format")
	if !govalidator.IsEmail(mail) {
		return err
	}
	return nil
}
