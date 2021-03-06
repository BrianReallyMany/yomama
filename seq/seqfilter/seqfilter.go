package seqfilter

import (
	. "github.com/BrianReallyMany/yomama/seq"
	"fmt"
)

type SeqFilterOptions struct {
	MinLength bool
	MinLengthValue int
	MaxLength bool
	MaxLengthValue int
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

func SeqPasses(seq *Seq, opts *SeqFilterOptions) bool {
	var pass bool
	if opts.MinLength {
		pass = minLengthTest(seq, opts)
		if !pass {
			return false
		}
	}
	if opts.MaxLength {
		pass = maxLengthTest(seq, opts)
		if !pass {
			return false
		}
	}
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
	if opts.NumberOfNs {
		pass = numberOfNsTest(seq, opts)
		if !pass {
			return false
		}
	}
	if opts.HomopolymerRun {
		pass = homopolymerRunTest(seq, opts)
		if !pass {
			return false
		}
	}
	return true
}

func minLengthTest(seq *Seq, opts *SeqFilterOptions) bool {
	if len(seq.Bases) < opts.MinLengthValue {
		return false
	}
	return true
}

func maxLengthTest(seq *Seq, opts *SeqFilterOptions) bool {
	if len(seq.Bases) > opts.MaxLengthValue {
		return false
	}
	return true
}

func avgValueTest(seq *Seq, opts *SeqFilterOptions) bool {
	avg := seq.AvgScore()
	if avg >= opts.MinAvgVal {
		return true
	}
	return false
}

func slidingWindowTest(seq *Seq, opts *SeqFilterOptions) bool {
	length := len(seq.Bases)
	winSize := opts.SlidingWinSize
	minScore := float32(opts.SlidingWinMinScore)
	for i := 0; i < length - winSize + 1; i++ {
		avg := avgValOfSlice(seq.Scores[i:i+winSize])
		if avg < minScore {
			return false
		}
	}
	return true
}

func numberOfNsTest(seq *Seq, opts *SeqFilterOptions) bool {
	count := 0
	bases := []byte(seq.Bases)
	for _, base := range bases {
		if base == 'N' || base == 'n' {
			count++
		}
	}
	if count > opts.MaxNumberOfNs {
		return false
	}
	return true
}

func homopolymerRunTest(seq *Seq, opts *SeqFilterOptions) bool {
	run := 1
	lastbase := byte('*')
	bases := []byte(seq.Bases)
	for _, base := range bases {
		if base == lastbase {
			run++
			if run > opts.MaxHomopolymerRun {
				return false
			}
		} else {
			run = 1
			lastbase = base
		}
	}
	return true
}

func avgValOfSlice(sl []int) float32 {
	total := 0
	for _, score := range sl {
		total += score
	}
	return float32(total) / float32(len(sl))
}
