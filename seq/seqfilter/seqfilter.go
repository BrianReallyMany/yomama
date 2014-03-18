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
	if opts.AvgVal {
		pass := avgValueTest(seq, opts)
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
