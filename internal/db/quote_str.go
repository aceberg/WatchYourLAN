package db

import "strings"

func quoteStr(str string) string {
	return strings.ReplaceAll(str, "'", "''")
}

func unquoteStr(str string) string {
	return strings.ReplaceAll(str, "''", "'")
}
