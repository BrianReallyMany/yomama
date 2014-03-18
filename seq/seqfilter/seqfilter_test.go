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

func getLousySeq() Seq {
	return Seq{"foo_seq", "GATTACA", "5 7 9 11 13 11 9", "", "", false}
}

func TestSeqPasses(t *testing.T) {
	opts := getAvgValueTestFilterOptions()
	seq := getLousySeq()
	ok := SeqPasses(seq, opts)
	if ok {
		t.Errorf("SeqPasses returned true; expected false")
	}
}

