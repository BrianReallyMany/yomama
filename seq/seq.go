package seq

import (
	"strings"
	"errors"
)

type Seq struct {
	Header  string
	Bases   string
	Scores  string
	Locus   string
	Sample  string
	Reverse bool
}

func TrimEnds(seq Seq, fromBegin, fromEnd int) (Seq, error) {
	// Check for error
	if len(seq.Bases) < fromBegin + fromEnd {
		return Seq{}, errors.New("TrimEnds called on seq that is too short")
	}
	// Trim Bases
	allBases := []byte(seq.Bases)
	bases := string(allBases[fromBegin:len(seq.Bases)-fromEnd])
	// Trim Scores
	allScores := strings.Split(seq.Scores, " ")
	scoreSlice := allScores[fromBegin:len(allScores)-fromEnd]
	scores := strings.Join(scoreSlice, " ")
	// Return new Seq with correct Bases and Scores
	return Seq{seq.Header, bases, scores, seq.Locus, seq.Sample, seq.Reverse}, nil
}
