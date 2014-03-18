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
	// Should pass with window size of 4
	if !ok {
		t.Errorf("SeqPasses returned false; expected true")
	}
	// Should fail with window size of 3
	opts.SlidingWinSize = 3
	ok = SeqPasses(seq, opts)
	if ok {
		t.Errorf("SeqPasses returned true; expected false")
	}
}

func TestAvgValOfSlice(t *testing.T) {
	testslice1 := []int{4, 5, 6, 7}
	avg1 := avgValOfSlice(testslice1)
	if avg1 != 5.5 {
		t.Errorf("avgValOfSlice returned %v; expected 5.5", avg1)
	}
	avg2 := avgValOfSlice(testslice1[1:])
	if avg2 != 6.0 {
		t.Errorf("avgValOfSlice returned %v; expected 6.0", avg2)
	}
}

func TestSeqPassesNumberOfNs(t *testing.T) {
	opts := &SeqFilterOptions{}
	opts.NumberOfNs = true
	opts.MaxNumberOfNs = 3
	seq := getLousySeq()
	ok := SeqPasses(seq, opts)
	// No Ns at all, should pass
	if !ok {
		t.Errorf("SeqPasses returned false, expected true")
	}
	seq.Bases = "GATTNNN"
	ok = SeqPasses(seq, opts)
	// 3 Ns, should still pass
	if !ok {
		t.Errorf("SeqPasses returned false, expected true")
	}
	seq.Bases = "GATNNNN"
	ok = SeqPasses(seq, opts)
	// 4 Ns, should fail
	if ok {
		t.Errorf("SeqPasses returned true, expected false")
	}

}

func TestSeqPassesHomopolymerRun(t *testing.T) {
	opts := &SeqFilterOptions{}
	opts.HomopolymerRun = true
	opts.MaxHomopolymerRun = 3
	seq := getLousySeq()
	ok := SeqPasses(seq, opts)
	// seq.Bases = "GATTACA", so should pass
	if !ok {
		t.Errorf("SeqPasses returned false, expected true")
	}
	seq.Bases = "GATTTCA"
	// 3 Ts in a row, should still pass
	ok = SeqPasses(seq, opts)
	if !ok {
		t.Errorf("SeqPasses returned false, expected true")
	}
	seq.Bases = "GATTTTA"
	// 4 Ts, should fail now
	ok = SeqPasses(seq, opts)
	if ok {
		t.Errorf("SeqPasses returned true, expected false")
	}
}
