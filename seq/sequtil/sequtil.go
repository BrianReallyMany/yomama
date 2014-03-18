package sequtil

// Returns the number of mismatches between an oligo pair and a raw sequence
// Second oligo string is reverse complemented before comparison
func MatchBeginAndEnd(oligoseqs [2]string, rawseq string) int {
	misses := 0
	frontoligo := oligoseqs[0]
	rearoligo := ReverseComplement(oligoseqs[1])
	beginraw := rawseq[:len(frontoligo)]
	endraw := rawseq[len(rawseq) - len(rearoligo):]
	// Note that it is required to pass the arguments to
	// NumberMismatches in this order
	misses += NumberMismatches(frontoligo, beginraw)
	misses += NumberMismatches(rearoligo, endraw)
	return misses
}

func NumberMismatches(oligoseq, rawseq string) int {
	count := 0
	len1 := len(oligoseq)
	len2 := len(rawseq)
	shorterLen := len1

	// No mismatches if oligo is empty string
	if len1 == 0 {
		return 0
	}

	// Penalty for seqs of unequal length
	if diff := len1 - len2; diff != 0 {
		if diff < 0 {
			count = -diff
		} else {
			shorterLen = len2
			count = diff
		}
	}

	for i := 0; i < shorterLen; i++ {
		// Note that it is required to provide the arguments to MatchBase
		// in this order.
		if !MatchBase(string(oligoseq[i]), string(rawseq[i])) {
			count++
		}
	}
	return count
}

// Copied from StackOverflow. Very clever.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Complement(s string) string {
	result := make([]rune, len(s))
	var compdict = map[rune]rune{
		'A': 'T',
		'C': 'G',
		'G': 'C',
		'T': 'A',
		'a': 't',
		'c': 'g',
		'g': 'c',
		't': 'a',
		'N': 'N',
		'n': 'n',
	}
	runes := []rune(s)
	for i, rune := range runes {
		result[i] = compdict[rune]
	}
	return string(result)
}


func ReverseComplement(seq string) string {
	reversed := Reverse(seq)
	return Complement(reversed)
}

