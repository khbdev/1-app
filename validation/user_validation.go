package validation

import (
	"1-loyiha/model"
	"errors"
	"strings"
)


func ValidationUser(user model.User) error {
	if strings.TrimSpace(user.Name) == "" {
		return errors.New("name is required")
	}
	if !strings.Contains(user.Email, "@") {
		return errors.New("Invalid email")
	}
	return nil
}