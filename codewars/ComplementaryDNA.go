package codewars

func DNAStrand(dna string) string {
	var newString string
	for i := 0; i < len(dna); i++ {
		switch dna[i] {
		case 'A':
			newString += "T"
		case 'T':
			newString += "A"
		case 'G':
			newString += "C"
		case 'C':
			newString += "G"
		}
	}
	return newString
}
