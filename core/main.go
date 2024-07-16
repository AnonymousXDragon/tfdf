package core

import "math"

func Tfdf(term string, doc []string, docs [][]string) float64 {
	tf := term_freq(term, doc)
	df := document_freq(term, docs)

	return tf * math.Log(1+df)
}
