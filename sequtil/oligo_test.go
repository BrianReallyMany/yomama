package sequtil

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
func TestOligoTextToMapsAndLinkerSlice(t *testing.T) {
	text := getFakeOligoText()
	primerMap, barcodeMap, linkers := OligoTextToMapsAndLinkerSlice(text)
	primerinput := [2]string{"GTGTAT", "ATCAAT"}
	primeroutput := primerMap[primerinput]

	if (primeroutput != "locus1") {
		t.Errorf("OligoTextToMap returned a primerMap with m[%s] = %s, wanted 'locus1'", primerinput, primeroutput)
	}
	barcodeinput := [2]string{"ATCGTACGTC", "TAGAATAAAC"}
	barcodeoutput := barcodeMap[barcodeinput]
	if (barcodeoutput != "sample1") {
		t.Errorf("OligoTextToMap returned a barcodeMap with m[%s] = %s, wanted 'sample1'", barcodeinput, barcodeoutput)
	}
	linkerinput := [2]string{"TCGGCAGCGTCAGAT", "GACTGTGGCAACACC"}
	linkeroutput := linkers[0]
	if (linkerinput != linkeroutput) {
		t.Errorf("OligoTextToMap returned linkers with [0] = %s, wanted '%s'", linkeroutput, linkerinput)
	}
}

func TestValidateOligosTextTrue(t *testing.T) {
	text := getFakeOligoText()
	if _, ok := ValidateOligosText(text); !ok {
		t.Errorf("ValidateOligosText failed on valid text:\n%s", text)
	}
}

func TestValidateOligosTextFalse(t *testing.T) {
	text := getInvalidOligosText()
	fmt.Printf("Unit test should print error message: ")
	if _, ok := ValidateOligosText(text); ok {
		t.Errorf("ValidateOligosText returned 'ok' on bad input:\n%s", text)
	}
}


func TestValidateOligosTextNumberOfLinkers(t *testing.T) {
	text := getFakeOligoText()
	if numLinkers, _ := ValidateOligosText(text); numLinkers != 1 {
		t.Errorf("ValidateOligosText returned numLinkers = %d, expected 1", numLinkers)
	}
}

func TestValidateOligoLine(t *testing.T) {
	line := "barcode\tATCGTACGTC\tTAGAATAAAC\tsample1\n"
	oligotype := ValidateOligoLine(line)
	if (oligotype != "barcode") {
		t.Errorf("ValidateOligoLine returned type=%s; wanted 'barcode'", oligotype)
	}
}

func TestValidateOligoBadLine(t *testing.T) {
	line := "acme_oligo_much_sequence\tGATTACA\tGATTACA\tmany_sample\n"
	fmt.Printf("Unit test should print error message: ")
	oligotype := ValidateOligoLine(line)
	if oligotype != "" {
		t.Errorf("ValidateOligoLine returned type=%s; expected empty string for invalid input", oligotype)
	}
}
