package sequtil

import "strings"

func FindFlankingSeqs(seq string, flankers [][]string) []string {
	for _, pair := range flankers {
		if (strings.HasPrefix(seq, pair[0]) && strings.HasSuffix(seq, pair[1])) {
			return pair
		}
	}
	return nil
}
