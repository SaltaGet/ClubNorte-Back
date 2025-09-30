package repositories

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

func IsDuplicateError(err error) bool {
	return errors.Is(err, gorm.ErrDuplicatedKey) ||
		strings.Contains(err.Error(), "Duplicate entry") ||
		strings.Contains(err.Error(), "duplicate key value violates unique constraint")
}

func DuplicateField(err error) string {
	msg := err.Error()
	switch {
	case strings.Contains(msg, "email"):
		return "email"
	case strings.Contains(msg, "username"):
		return "username"
	default:
		return ""
	}
}
