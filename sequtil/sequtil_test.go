package sequtil

import "testing"

func makeTestMap() map[[2]string]string {
	m := make(map[[2]string]string)
	keys1 := [2]string{"GATT", "GAT"}
	keys2 := [2]string{"CACA", "GAGA"}
	keys3 := [2]string{"GATT", "ACAG"}
	m[keys1] = "locus1"
	m[keys2] = "locus2"
	m[keys3] = "locus3"
	return m
}

// yo dawg
func TestMakeTestMap(t *testing.T) {
	var m map[[2]string]string
	m = makeTestMap()
	expected := "locus2"
	in := [2]string{"CACA", "GAGA"}
	x := m[in]
	if (x != expected) {
		t.Errorf("makeTestMap[%s] = %s, want %s", in, x, expected)
	}
}

func TestMatchBeginAndEndTrue(t *testing.T) {
	testraw := "GATTACA"
	testoligos := [2]string{"GAT", "ACA"}
	match := MatchBeginAndEnd(testoligos, testraw, 0)
	expected := true
	if (match != expected) {
		t.Errorf("MatchBeginAndEnd(%s, %s, 0) = %t, want %t", testoligos, testraw, match, expected)
	}
}

func TestMatchBeginAndEndFalse(t *testing.T) {
	testraw := "GATTACA"
	testoligos := [2]string{"GGG", "GGG"}
	match := MatchBeginAndEnd(testoligos, testraw, 0)
	if (match) {
		t.Errorf("MatchBeginAndEnd(%s, %s, 0) returned true.", testoligos, testraw)
	}
}

func TestMatchBeginAndEndMistakesAllowedTrue(t *testing.T) {
	testraw := "GATTACA"
	testoligos := [2]string{"GAA", "ACA"}
	match := MatchBeginAndEnd(testoligos, testraw, 1)
	if (!match) {
		t.Errorf("MatchBeginAndEnd(%s, %s, 1) returned false.", testoligos, testraw)
	}
}
	

func TestMatchBeginAndEndMistakesAllowedFalse(t *testing.T) {
	testraw := "GATTACA"
	testoligos := [2]string{"GCC", "ACA"}
	match := MatchBeginAndEnd(testoligos, testraw, 1)
	if (match) {
		t.Errorf("MatchBeginAndEnd(%s, %s, 1) returned true.", testoligos, testraw)
	}
}

func TestNumberMismatches(t *testing.T) {
	seq1 := "GGG"
	seq2 := "GGA"
	mismatches := NumberMismatches(seq1, seq2)
	if (mismatches != 1) {
		t.Errorf("NumberMismatches(%s, %s) returned %d, want 1", seq1, seq2, mismatches)
	}
}

func TestNumberMismatchesZero(t *testing.T) {
	seq1 := "GGA"
	seq2 := "GGA"
	mismatches := NumberMismatches(seq1, seq2)
	if (mismatches != 0) {
		t.Errorf("NumberMismatches(%s, %s) returned %d, want 0", seq1, seq2, mismatches)
	}
}

func TestNumberMismatchesUnequalLengths(t *testing.T) {
	seq1 := "GGGCC"
	seq2 := "GGA"
	mismatches := NumberMismatches(seq1, seq2)
	if (mismatches != 3) {
		t.Errorf("NumberMismatches(%s, %s) returned %d, want 3", seq1, seq2, mismatches)
	}
}
