package sequtil

import (
	"testing"
	"bytes"
)

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
	testoligos := [2]string{"GAT", "TGT"}
	mismatches := MatchBeginAndEnd(testoligos, testraw, 8)
	expected := 0
	if (mismatches != expected) {
		t.Errorf("MatchBeginAndEnd(%s, %s, 8) = %d, want %d", testoligos, testraw, mismatches, expected)
	}
}

func BenchmarkMatchBeginAndEnd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		oligos := [2]string{"ACTACTATGT", ""}
		rawseq := "ACTACTATGTCGTAGACTAGTACGGTAGCAGAGACTTGGTCTCGCCACAACGCAACAAAGCAGACACTTTACAGCAACAAGAAGGCAAAGCAAGATGAAATGAAAAATGTCTACGCATACATACATACATACATACATACATAAGTAAACATAGTAATAAACCGTAAATGGTTGTGTGTATGTGCTGTAGAACCATGTCGTCAGTGT"
		MatchBeginAndEnd(oligos, rawseq, 20)
	}
}

func TestNumberMismatches(t *testing.T) {
	seq1 := []byte("ggg")
	seq2 := []byte("gga")
	mismatches := NumberMismatches(seq1, seq2, 4)
	if (mismatches != 1) {
		t.Errorf("NumberMismatches(%s, %s, 4) returned %d, want 1", seq1, seq2, mismatches)
	}
}

func BenchmarkNumberMismatches(b *testing.B) {
	for i := 0; i < b.N; i++ {
		seq1 := bytes.ToLower([]byte("AGTACCGCCCTGTTCTAACCCTTAAACGATGCCCAGCTGCAATTTGGGGTGTA"))
		seq2 := bytes.ToLower([]byte("AGTACCGCTCTGTTCTAACCCTAAAACGATGCCCGGCTGCAATTTGTGGTGTA"))
		NumberMismatches(seq1, seq2, 100)
	}
}

func BenchmarkNumberMismatchesNoDifferences(b *testing.B) {
	for i := 0; i < b.N; i++ {
		seq1 := []byte("AGTACCGCCCTGTTCTAACCCTTAAACGATGCCCAGCTGCAATTTGGGGTGTA")
		seq2 := []byte("AGTACCGCCCTGTTCTAACCCTTAAACGATGCCCAGCTGCAATTTGGGGTGTA")
		NumberMismatches(seq1, seq2, 100)
	}
}

func TestNumberMismatchesZero(t *testing.T) {
	seq1 := []byte("gga")
	seq2 := []byte("gga")
	mismatches := NumberMismatches(seq1, seq2, 4)
	if (mismatches != 0) {
		t.Errorf("NumberMismatches(%s, %s, 4) returned %d, want 0", seq1, seq2, mismatches)
	}
}

func TestNumberMismatchesOligoStringEmpty(t *testing.T) {
	oligo := []byte("")
	raw := []byte("ggg")
	mismatches := NumberMismatches(oligo, raw, 4)
	if mismatches != 0 {
		t.Errorf("NumberMismatches('', 'GGG', 4) returned %s, want 0", mismatches)
	}
}

func TestNumberMismatchesUnequalLengths(t *testing.T) {
	seq1 := []byte("gggcc")
	seq2 := []byte("gga")
	mismatches := NumberMismatches(seq1, seq2, 6)
	if (mismatches != 3) {
		t.Errorf("NumberMismatches(%s, %s, 6) returned %d, want 3", seq1, seq2, mismatches)
	}
}

func TestNumberMismatchesExceedsMax(t *testing.T) {
	seq1 := []byte("gggcc")
	seq2 := []byte("ggaaa")
	mismatches := NumberMismatches(seq1, seq2, 1)
	if (mismatches != 1) {
		t.Errorf("NumberMismatches(%s, %s, 1) returned %d, want 1", seq1, seq2, mismatches)
	}
}

func TestReverse(t *testing.T) {
	if result := Reverse("abc123"); result != "321cba" {
		t.Errorf("Reverse('abc123') = %s, expected '321cba'", result)
	}
}

func TestComplement(t *testing.T) {
	if result := Complement("ACGT"); result != "TGCA" {
		t.Errorf("Complement('ACGT') = %s; expected 'TGCA'", result)
	}
}

func TestReverseComplementOneBase(t *testing.T) {
	inseq := "C"
	outseq := ReverseComplement(inseq)
	if outseq != "G" {
		t.Errorf("ReverseComplement('C') returned %s, expected 'G'", outseq)
	}
}

func TestReverseComplementSeq(t *testing.T) {
	inseq := "CAT"
	outseq := ReverseComplement(inseq)
	if outseq != "ATG" {
		t.Errorf("ReverseComplement('CAT') returned %s, expected 'ATG'", outseq)
	}
}

func TestQualStringToIntSlice(t *testing.T) {
	qualstring := "20 30 25 40"
	intslice := QualStringToIntSlice(qualstring)
	expected := []int{20, 30, 25, 40}
	for i, score := range intslice {
		if score != expected[i] {
			t.Errorf("QualStringToIntSlice returned %d; expected %d", score, expected[i])
		}
	}
}

func TestStringToPhredScoreSlice(t *testing.T) {
	string1 := "A"
	scoreslice1phred33 := StringToPhredScoreSlice(string1, false)
	if scoreslice1phred33[0] != 32 {
		t.Errorf("StringToPhredScoreSlice returned %d, expected 32", scoreslice1phred33[0])
	}
	string2 := "!\"#$%&'()*+,-./0"
	expected2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	scoreslice2phred33 := StringToPhredScoreSlice(string2, false)
	for i, score := range scoreslice2phred33 {
		if score != expected2[i] {
			t.Errorf("StringToPhredScoreSlice returned %d, expected %d", score, expected2[i])
		}
	}
	string3 := "@Jh"
	expected3 := []int{0, 10, 40}
	scoreslice3phred64 := StringToPhredScoreSlice(string3, true)
	for i, score := range scoreslice3phred64 {
		if score != expected3[i] {
			t.Errorf("StringToPhredScoreSlice returned %d, expected %d", score, expected3[i])
		}
	}
}

