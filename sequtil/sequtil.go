package sequtil

func MatchBeginAndEnd(flankseqs [2]string, seq string, mismatchesAllowed int) bool {
	// need func to count mismatches between 2 seqs
	misses := 0
	frontflank := flankseqs[0]
	rearflank := flankseqs[1]
	beginseq := seq[:len(frontflank)]
	endseq := seq[len(seq) - len(rearflank):]
	// Note that it is required to pass the arguments to
	// NumberMismatches in this order
	misses += NumberMismatches(frontflank, beginseq)
	misses += NumberMismatches(rearflank, endseq)
	if misses > mismatchesAllowed {
		return false
	}
	return true 	// TODO
}

func NumberMismatches(utilityseq, rawseq string) int {
	count := 0
	len1 := len(utilityseq)
	len2 := len(rawseq)
	shorterLen := len1

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
		if !MatchBase(string(utilityseq[i]), string(rawseq[i])) {
			count++
		}
	}
	return count
}
