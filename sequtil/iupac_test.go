package sequtil

import "testing"

func TestMatchBaseTrivialTrue(t *testing.T) {
	if (!MatchBase("A", "A")) {
		t.Errorf("MatchBase('A', 'A') returned false.")
	}
}

func TestMatchBaseTrivialFalse(t *testing.T) {
	if (MatchBase("A", "C")) {
		t.Errorf("MatchBase('A', 'C') returned true.")
	}
}

func TestMatchBaseAtoRFalse(t *testing.T) {
	if (MatchBase("A", "R")) {
		t.Errorf("MatchBase('A', 'R') returned true. Shouldn't have 'R' in raw sequence.")
	}
}

func TestMatchBaseHtoGFalse(t *testing.T) {
	if (MatchBase("H", "G")) {
		t.Errorf("MatchBase('H', 'G') returned true.")
	}
}

func TestMatchBaseCtoNTrue(t *testing.T) {
	if (!MatchBase("C", "N")) {
		t.Errorf("MatchBase('C', 'N') returned false.")
	}
}

func TestMatchBaseStoTFalse(t *testing.T) {
	if (MatchBase("S", "T")) {
		t.Errorf("MatchBase('S', 'T') returned true.")
	}
}

func TestMatchBaseatoATrue(t *testing.T) {
	if (!MatchBase("a", "A")) {
		t.Errorf("MatchBase('a', 'A') returned false.")
	}
}
