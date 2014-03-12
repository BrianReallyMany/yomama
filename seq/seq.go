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

func (s *Seq) TrimEnds(fromBegin, fromEnd int) error {
	// Check for error
	if len(s.Bases) < fromBegin + fromEnd {
		return errors.New("TrimEnds called on seq that is too short")
	}
	// Trim Bases
	allBases := []byte(s.Bases)
	s.Bases = string(allBases[fromBegin:len(s.Bases)-fromEnd])
	// Trim Scores
	allScores := strings.Split(s.Scores, " ")
	scoreSlice := allScores[fromBegin:len(allScores)-fromEnd]
	s.Scores = strings.Join(scoreSlice, " ")
	return nil
}
