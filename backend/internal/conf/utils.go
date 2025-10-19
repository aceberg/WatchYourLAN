package conf

import "strings"

// stripQuotes removes surrounding quotes from a string value
func stripQuotes(value string) string {
	if len(value) >= 2 {
		if (strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`)) ||
			(strings.HasPrefix(value, `'`) && strings.HasSuffix(value, `'`)) {
			return value[1 : len(value)-1]
		}
	}
	return value
}