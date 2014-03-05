package oligo

import (
	"fmt"
	"testing"
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

func TestValidateOligoTextTrue(t *testing.T) {
	text := getFakeOligoText()
	if ok := ValidateOligoText(text); !ok {
		t.Errorf("ValidateOligoText failed on valid text:\n%s", text)
	}
}

func TestValidateOligoTextFalse(t *testing.T) {
	text := getInvalidOligoText()
	if ok := ValidateOligoText(text); ok {
		t.Errorf("ValidateOligoText returned 'ok' on bad input:\n%s", text)
	}
}

func TestCountLinkers(t *testing.T) {
	text := getFakeOligoText()
	if numLinkers := CountLinkers(text); numLinkers != 1 {
		t.Errorf("ValidateOligoText returned numLinkers = %d, expected 1", numLinkers)
	}
}

func TestValidateOligoLineGoodLine(t *testing.T) {
	line := "barcode\tATCGTACGTC\tTAGAATAAAC\tsample1\n"
	ok := ValidateOligoLine(line)
	if !ok {
		t.Errorf("ValidateOligoLine returned not ok; wanted ok")
	}
}

func TestValidateOligoBadLine(t *testing.T) {
	line := "very_oligo_much_sequence\tGATTACA\tGATTACA\tmany_sample\n"
	fmt.Printf("(Unit test should print error message:) ")
	ok := ValidateOligoLine(line)
	if ok {
		t.Errorf("ValidateOligoLine returned ok; expected not ok")
	}
}
