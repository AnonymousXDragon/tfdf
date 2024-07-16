package core

import "strings"

func Tokenize(doc string) []string {
	return strings.Fields(strings.ToLower(doc))
}
