package sortseq

import (
	"testing"
	. "github.com/BrianReallyMany/yomama/seq"
)

func getFakeOligoText() string {
	result := "barcode\tATCGTACGTC\tTAGAATAAAC\tsample1\n"
	result += "linker\tTCGGCAGCGTCAGAT\tGACTGTGGCAACACC\n"
	result += "primer\tGTGTAT\tATCAAT\tlocus1\n"
	return result
}

func getInvalidOligoText() string {
	result := "barcode\tATCGTACGTC\tTAGAATAAAC\tsample1\n"
	result += "linker\tTCGGCAGCGTCAGAT\tGACTGTGGCAACACC\n"
	result += "primer\tGTGTAT\tATCAAT\t\n"	//missing primer id value
	return result
}

func getSeqSorter() SeqSorter{
	var priMap = map[[2]string]string{
		[2]string{"GTGTAT", "ATCAAT"}: "locus1",
	}
	var barMap = map[[2]string]string{
		[2]string{"ATCGTACGTC", "TAGAATAAAC"}: "sample1",
	}
	var links = [][2]string{[2]string{"TCGGCAGCGTCAGAT", "GACTGTGGCAACACC"}}
	return SeqSorter{priMap, barMap, links}
}

func TestNewSeqSorterValidInput(t *testing.T) {
	text := getFakeOligoText()
	sorter, _ := NewSeqSorter(text)
	primerinput := [2]string{"GTGTAT", "ATCAAT"}
	primeroutput := sorter.primerMap[primerinput]

	if (primeroutput != "locus1") {
		t.Errorf("OligoTextToMap returned a primerMap with m[%s] = %s, wanted 'locus1'", primerinput, primeroutput)
	}
	barcodeinput := [2]string{"ATCGTACGTC", "TAGAATAAAC"}
	barcodeoutput := sorter.barcodeMap[barcodeinput]
	if (barcodeoutput != "sample1") {
		t.Errorf("OligoTextToMap returned a barcodeMap with m[%s] = %s, wanted 'sample1'", barcodeinput, barcodeoutput)
	}
	linkerinput := [2]string{"TCGGCAGCGTCAGAT", "GACTGTGGCAACACC"}
	linkeroutput := sorter.linkers[0]
	if (linkerinput != linkeroutput) {
		t.Errorf("OligoTextToMap returned linkers with [0] = %s, wanted '%s'", linkeroutput, linkerinput)
	}
}

func TestNewSeqSorterInvalidInput(t *testing.T) {
	text := getInvalidOligoText()
	_, err := NewSeqSorter(text)
	if err == nil {
		t.Errorf("NewSeqSorter failed to return error on invalid input, expected error.")
	}
}

func TestSortSeq(t *testing.T) {
	sorter := getSeqSorter()
	seq := Seq{"foo_seq", "ATCGTACGTCTCGGCAGCGTCAGATGTGTATgattacaATTGATGGTGTTGCCACAGTCGTTTATTCTA", "40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40"}
	if sorter.barcodeMap[[2]string{"ATCGTACGTC", "TAGAATAAAC"}] != "sample1" {
		t.Errorf("foo error.")
	}
	sorted := sorter.SortSeq(seq)
	if bases := sorted.seq.Bases; bases != "gattaca" {
		t.Errorf("SortSeq returned a SortedSeq with bases = %s; expected 'gattaca'", bases)
	}
}
