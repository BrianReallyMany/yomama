package seq

import "testing"

func makeTestSeq() Seq {
	header := "foo_seq"
	bases := "GATTACA"
	scores := "30 30 30 30 30 30 30"
	locus := "locus1"
	sample := "sample1"
	reverse := false
	return Seq{header, bases, scores, locus, sample, reverse}
}


func TestReverseComplementSeq(t *testing.T) {
	seq := makeTestSeq()
	if seq.Bases != "GATTACA" {
		t.Errorf("makeTestSeq returned seq with Bases %s; expected 'GATTACA'", seq.Bases)
	}
}

func TestTrimEnds(t *testing.T) {
	seq := makeTestSeq()
	if len(seq.Bases) != 7 {
		t.Errorf("makeTestSeq gave me a seq with %d bases, expected 7...", len(seq.Bases))
	}
	newSeq, err := TrimEnds(seq, 2, 2)
	if err != nil {
		t.Errorf("TrimEnds returned an error...")
	}
	if newSeq.Bases != "TTA" {
		t.Errorf("TrimEnds returned bases = %s, expected 'TTA'", newSeq.Bases)
	}
	if newSeq.Scores != "30 30 30" {
		t.Errorf("TrimEnds returned scores = %s, expected '30 30 30'", newSeq.Scores)
	}
}

func TestTrimEndsError(t *testing.T) {
	seq := makeTestSeq()
	_, err := TrimEnds(seq, 4, 4)
	if err == nil {
		t.Errorf("TrimEnds did not return an error; it should have.")
	}
}
