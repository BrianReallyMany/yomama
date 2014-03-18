package seq

import (
	"strings"
	"strconv"
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

func (s *Seq) ToString() string {
	result := "Header: " + s.Header + "\n"
	result += "Bases: " + s.Bases + "\n"
	result += "Scores: " + s.Scores + "\n"
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
	allScores := strings.Split(s.Scores, " ")
	scoreSlice := allScores[fromBegin:len(allScores)-fromEnd]
	s.Scores = strings.Join(scoreSlice, " ")
	return nil
}

func (s *Seq) AvgScore() float32 {
	scoreslice := s.ScoresAsSliceOfInts()
	total := 0
	for _, score := range scoreslice {
		total += score
	}
	return float32(total) / float32(len(scoreslice))
}

func (s *Seq) ScoresAsSliceOfInts() []int {
	scoreslice := make([]int, len(s.Bases))
	splitscores := strings.Split(s.Scores, " ")
	for i, score := range splitscores {
		intscore, err := strconv.Atoi(score)
		if err != nil {
			return scoreslice
		}
		scoreslice[i] = intscore
	}
	return scoreslice
}

