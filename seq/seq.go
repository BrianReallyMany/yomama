package seq

import (
	"strings"
	"strconv"
	"errors"
)

type Seq struct {
	Header  string
	Bases   string
	Scores  []int
	Locus   string
	Sample  string
	Reverse bool
}

func (s *Seq) ToString() string {
	result := "Header: " + s.Header + "\n"
	result += "Bases: " + s.Bases + "\n"
	result += "Scores: " + s.ScoresToString() + "\n"
	result += "Locus: " + s.Locus + "\n"
	result += "Sample: " + s.Sample + "\n"
	result += "Reverse: " + strconv.FormatBool(s.Reverse)
	return result
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
	s.Scores = s.Scores[fromBegin:len(s.Scores)-fromEnd] 
	return nil
}

func (s *Seq) AvgScore() float32 {
	total := 0
	for _, score := range s.Scores {
		total += score
	}
	return float32(total) / float32(len(s.Scores))
}

func (s *Seq) ScoresToString() string {
	scorestring := ""
	for _, score := range s.Scores {
		scorestring += strconv.Itoa(score) + " "
	}
	return strings.Trim(scorestring, " ")
}
