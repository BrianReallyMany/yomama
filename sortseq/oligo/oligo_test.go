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

func getInvalidOligosText() string {
	result := "barcode\tATCGTACGTC\tTAGAATAAAC\tsample1\n"
	result += "linker\tTCGGCAGCGTCAGAT\tGACTGTGGCAACACC\n"
	result += "primer\tGTGTAT\tATCAAT\t\n"	//missing primer id value
	return result
}
func TestOligoTextToSeqSorter(t *testing.T) {
	text := getFakeOligoText()
	sorter, _ := OligoTextToSeqSorter(text)
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

func TestValidateOligosTextTrue(t *testing.T) {
	text := getFakeOligoText()
	if ok := ValidateOligosText(text); !ok {
		t.Errorf("ValidateOligosText failed on valid text:\n%s", text)
	}
}

func TestValidateOligosTextFalse(t *testing.T) {
	text := getInvalidOligosText()
	fmt.Printf("Unit test should print error message: ")
	if ok := ValidateOligosText(text); ok {
		t.Errorf("ValidateOligosText returned 'ok' on bad input:\n%s", text)
	}
}

func TestCountLinkers(t *testing.T) {
	text := getFakeOligoText()
	if numLinkers := CountLinkers(text); numLinkers != 1 {
		t.Errorf("ValidateOligosText returned numLinkers = %d, expected 1", numLinkers)
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
	line := "acme_oligo_much_sequence\tGATTACA\tGATTACA\tmany_sample\n"
	fmt.Printf("Unit test should print error message: ")
	ok := ValidateOligoLine(line)
	if ok {
		t.Errorf("ValidateOligoLine returned ok; expected not ok")
	}
}
