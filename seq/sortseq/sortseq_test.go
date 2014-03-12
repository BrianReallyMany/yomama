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
	return SeqSorter{priMap, barMap, links, SeqSorterOptions{}}
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
	seq := Seq{"foo_seq", "ATCGTACGTCTCGGCAGCGTCAGATGTGTATgattacaATTGATGGTGTTGCCACAGTCGTTTATTCTA", "40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40 40", "", "", true}
	sorted := sorter.SortSeq(seq)
	if bases := sorted.Bases; bases != "gattaca" {
		t.Errorf("SortSeq returned a Seq with bases = %s; expected 'gattaca'", bases)
	}
}

func TestBestMatch(t *testing.T) {
	testOligos := make([][2]string, 2)
	testOligos[0] = [2]string{"GTGTAA", "ATCAAT"}
	testOligos[1] = [2]string{"GTGTAT", "ATCAAT"}
	best, num := bestMatch(testOligos, "ATGTAAATTGAT") // 1 mismatch with testOligos[0], 2 with testOligos[1]
	if num != 1 {
		t.Errorf("bestMatch returned %d errors; expected 1.", num)
	}
	if best[0] != "GTGTAA" {
		t.Errorf("bestMatch returned %s, expected 'GTGTAA'.", best[0])
	}
}

func TestGetSliceOfKeys(t *testing.T) {
	testMap := make(map[[2]string]string)
	testMap[[2]string{"foo", "bar"}] = "baz"
	testMap[[2]string{"dog", "cat"}] = "pig"
	keys := getSliceOfKeys(testMap)
	if len(keys) != 2 {
		t.Errorf("getSliceOfKeys returned slice of length %d; expected 2.", len(keys))
	}
	if !(keys[0][0] == "foo" || keys[1][0] == "foo") {
		t.Errorf("getSliceOfKeys returned slice without 'foo', expected foo. wtf.")
	}
}

