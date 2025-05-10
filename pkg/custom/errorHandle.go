package custom

import (
	"strings"
)

func IsUniqueConstraintError(err error, constraintName string) bool {
	if err == nil {
		return false
	}

	if strings.Contains(err.Error(), "23505") && (constraintName == "" || strings.Contains(err.Error(), constraintName)) {
		return true
	}

	return false
}
