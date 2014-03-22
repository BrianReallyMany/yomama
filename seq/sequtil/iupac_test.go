package sequtil

import "testing"

func TestMatchBaseTrivialTrue(t *testing.T) {
	if (!MatchBase('a', 'a')) {
		t.Errorf("MatchBase('a', 'a') returned false.")
	}
}

func TestMatchBaseTrivialFalse(t *testing.T) {
	if (MatchBase('a', 'c')) {
		t.Errorf("MatchBase('a', 'c') returned true.")
	}
}

//func TestMatchBaseAtoRFalse(t *testing.T) {
	//if (MatchBase("A", "R")) {
		//t.Errorf("MatchBase('A', 'R') returned true. Shouldn't have 'R' in raw sequence.")
	//}
//}
//
//func TestMatchBaseHtoGFalse(t *testing.T) {
	//if (MatchBase("H", "G")) {
		//t.Errorf("MatchBase('H', 'G') returned true.")
	//}
//}
//
//func TestMatchBaseCtoNFalse(t *testing.T) {
	//if (MatchBase("C", "N")) {
		//t.Errorf("MatchBase('C', 'N') returned true.")
	//}
//}
//
//func TestMatchBaseStoTFalse(t *testing.T) {
	//if (MatchBase("S", "T")) {
		//t.Errorf("MatchBase('S', 'T') returned true.")
	//}
//}

//func TestMatchBaseatoATrue(t *testing.T) {
	//if (!MatchBase("a", "A")) {
		//t.Errorf("MatchBase('a', 'A') returned false.")
	//}
//}

func BenchmarkMatchBase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MatchBase('A', 'C')
	}
}
