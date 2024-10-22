package application

import (
	"authorization_service/internal/core/cerror"
	"regexp"
)

func checkUsername(username string) error {
	regex, err := regexp.Compile("^[a-zA-Z0-9]+(_?[a-zA-Z0-9]+)*$")
	if err != nil {
		return err
	}

	if !regex.Match([]byte(username)) {
		return cerror.New(cerror.USERNAME_INVALID, "username is not valid")
	}

	return nil
}
