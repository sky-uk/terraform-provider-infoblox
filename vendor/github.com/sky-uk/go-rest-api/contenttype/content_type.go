package contenttype

import (
	"strings"
)

// GetType - parses a valid Content-Type string and returns
// the type found
// A valid Content-Type string matches this BNF:
// Content-Type := type "/" subtype *[";" parameter]
func GetType(content string) string {
	subs := strings.Split(content, ";")
	types := strings.Split(subs[0], "/")
	if len(types) > 1 {
		return strings.ToLower(types[1])
	}
	return strings.ToLower(types[0])
}
