package sequtil

func MatchBeginAndEnd(oligoseqs [2]string, rawseq string, mismatchesAllowed int) bool {
	// need func to count mismatches between 2 seqs
	misses := 0
	frontoligo := oligoseqs[0]
	rearoligo := oligoseqs[1]
	beginraw := rawseq[:len(frontoligo)]
	endraw := rawseq[len(rawseq) - len(rearoligo):]
	// Note that it is required to pass the arguments to
	// NumberMismatches in this order
	misses += NumberMismatches(frontoligo, beginraw)
	misses += NumberMismatches(rearoligo, endraw)
	if misses > mismatchesAllowed {
		return false
	}
	return true 	// TODO
}

func NumberMismatches(oligoseq, rawseq string) int {
	count := 0
	len1 := len(oligoseq)
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
		if !MatchBase(string(oligoseq[i]), string(rawseq[i])) {
			count++
		}
	}
	return count
}
