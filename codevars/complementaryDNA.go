package kata

func DNAStrand(dna string) string {
	var str string
	for _, letter := range dna {
		switch letter {
		case 'A':
			str += "T"
		case 'T':
			str += "A"
		case 'C':
			str += "G"
		case 'G':
			str += "C"
		}
	}
	return str
}
