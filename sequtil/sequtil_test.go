package sequtil

import "testing"

func TestFindFlankingSeqs(t *testing.T) {
	testPrimers0 := []string{"GATT", "GAT"}
	testPrimers1 := []string{"CACA", "GAGA"}
	testPrimers2 := []string{"GATT", "ACAG"}
	testPrimers := [][]string{testPrimers0, testPrimers1, testPrimers2}
	in := "GATTACAGATTACAG"
	var out = testPrimers2
	x := FindFlankingSeqs(in, testPrimers)
	if ((x[0] != out[0]) || (x[1] != out[1])) {
		t.Errorf("FindFlankingSeqs(%s) = %s, want %s", in, x, out)
	}
}

	
