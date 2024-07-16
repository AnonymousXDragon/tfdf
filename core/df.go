package core

func document_freq(term string, docs [][]string) float64 {
	var count float64

	for _, doc := range docs {
		for _, t := range doc {
			if t == term {
				count++
				break
			}
		}
	}

	return count
}
