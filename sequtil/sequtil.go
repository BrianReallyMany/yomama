package sequtil

import (
	"strings"
)

func FindFlankingSeqs(seq string, flankers [][]string) []string {
	for _, pair := range flankers {
		if (strings.HasPrefix(seq, pair[0]) && strings.HasSuffix(seq, pair[1])) {
			return pair
		}
	}
	return nil
}

func MatchBeginAndEnd(flankseqs [2]string, seq string, mismatchesAllowed int) bool {
	// need func to count mismatches between 2 seqs
	misses := 0
	frontflank := flankseqs[0]
	rearflank := flankseqs[1]
	beginseq := seq[:len(frontflank)]
	endseq := seq[len(seq) - len(rearflank):]
	misses += NumberMismatches(frontflank, beginseq)
	misses += NumberMismatches(rearflank, endseq)
	if misses > mismatchesAllowed {
		return false
	}
	return true 	// TODO
}

func NumberMismatches(seq1, seq2 string) int {
	count := 0
	len1 := len(seq1)
	len2 := len(seq2)
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
		if !MatchBase(string(seq1[i]), string(seq2[i])) {
			count++
		}
	}
	return count
}
