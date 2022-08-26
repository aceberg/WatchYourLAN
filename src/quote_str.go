package main

import "strings"

func quote_str(str string) string {
	return strings.ReplaceAll(str, "'", "''")
}

func unquote_str(str string) string {
	return strings.ReplaceAll(str, "''", "'")
}