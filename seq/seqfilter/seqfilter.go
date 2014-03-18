package seqfilter

import (
	. "github.com/BrianReallyMany/yomama/seq"
	"fmt"
)

type SeqFilterOptions struct {
	AvgVal bool
	MinAvgVal float32
	SlidingWin bool
	SlidingWinSize int
	SlidingWinMinScore int
	NumberOfNs bool
	MaxNumberOfNs int
	HomopolymerRun bool
	MaxHomopolymerRun int
}

type SeqFilterError struct {
	Filter string
	Problem string
}

func (e *SeqFilterError) Error() string {
	return fmt.Sprintf("SeqFilter error occurred. Filter=%s, Problem=%s", e.Filter, e.Problem)
}

func SeqPasses(seq Seq, opts *SeqFilterOptions) bool {
	var pass bool
	if opts.AvgVal {
		pass = avgValueTest(seq, opts)
		if !pass {
			return false
		}
	} 
	if opts.SlidingWin {
		pass = slidingWindowTest(seq, opts)
		if !pass {
			return false
		}
	}
	return true
}

func avgValueTest(seq Seq, opts *SeqFilterOptions) bool {
	avg := seq.AvgScore()
	if avg >= opts.MinAvgVal {
		return true
	}
	return false
}

func slidingWindowTest(seq Seq, opts *SeqFilterOptions) bool {
	//scoreslice := seq.ScoresAsSliceOfInts()
	return true
}

func avgValOfSlice(sl []int) float32 {
	total := 0
	for _, score := range sl {
		total += score
	}
	return float32(total) / float32(len(sl))
}
