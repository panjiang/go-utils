package sqls

import (
	"strings"
)

// IsDuplicateKeyError 键重复
func IsDuplicateKeyError(err error) bool {
	return strings.HasPrefix(err.Error(), "Error 1062:")
}
