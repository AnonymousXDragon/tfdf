package core

func term_freq(term string, doc []string) float64 {
	var count float64

	for _, t := range doc {
		if t == term {
			count++
		}
	}

	return count / float64(len(doc))
}
