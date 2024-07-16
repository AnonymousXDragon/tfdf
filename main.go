package main

import (
	"fmt"
	"tfdf/core"
)

func main() {
	docs := []string{
		"ronalado is a beast",
		"the complete player is ronaldo",
		"messi is a dribbler",
		"mbappe is the future",
	}

	var tdocs [][]string

	for _, doc := range docs {
		tdocs = append(tdocs, core.Tokenize(doc))
	}

	terms := []string{"ronaldo", "is", "messi", "mbappe"}

	for _, term := range terms {
		for i, doc := range tdocs {
			fmt.Printf("[%d]: %s: %f\n", i+1, term, core.Tfdf(term, doc, tdocs))
		}
	}
}
