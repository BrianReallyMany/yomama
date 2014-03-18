package seqfilter

import (
	"testing"
	. "github.com/BrianReallyMany/yomama/seq"
)

func getAvgValueTestFilterOptions() *SeqFilterOptions {
	opts := SeqFilterOptions{}
	opts.AvgVal = true
	opts.MinAvgVal = 10
	return &opts
}

func getSlidingWindowTestFilterOptions() *SeqFilterOptions {
	opts := SeqFilterOptions{}
	opts.SlidingWin = true
	opts.SlidingWinSize = 4
	opts.SlidingWinMinScore = 8
	return &opts
}

func getLousySeq() Seq {
	return Seq{"foo_seq", "GATTACA", "5 7 9 11 13 11 7", "", "", false}
}

func TestSeqPassesAvgVal(t *testing.T) {
	opts := getAvgValueTestFilterOptions()
	seq := getLousySeq()
	ok := SeqPasses(seq, opts)
	if ok {
		t.Errorf("SeqPasses returned true; expected false")
	}
	opts.MinAvgVal = 9
	ok = SeqPasses(seq, opts)
	if !ok {
		t.Errorf("SeqPasses returned false; expected true")
	}
}

func TestSeqPassesSlidingWindow(t *testing.T) {
	opts := getSlidingWindowTestFilterOptions()
	seq := getLousySeq()
	ok := SeqPasses(seq, opts)
	if !ok {
		t.Errorf("SeqPasses returned false; expected true")
	}
}

func TestAvgValOfSlice(t *testing.T) {
	testslice1 := []int{4, 5, 6, 7}
	avg1 := avgValOfSlice(testslice1)
	if avg1 != 5.5 {
		t.Errorf("avgValOfSlice returned %v; expected 5.5", avg1)
	}
}
