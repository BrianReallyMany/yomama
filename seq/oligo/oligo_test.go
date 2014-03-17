package oligo

import (
	"bufio"
	"strings"
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

func TestValidateOligoTextTrue(t *testing.T) {
	text := getFakeOligoText()
	buffer := bufio.NewReader(strings.NewReader(text))
	if ok := ValidateOligoText(buffer); !ok {
		t.Errorf("ValidateOligoText failed on valid text:\n%s", text)
	}
}

func TestValidateOligoTextFalse(t *testing.T) {
	text := getInvalidOligoText()
	buffer := bufio.NewReader(strings.NewReader(text))
	if ok := ValidateOligoText(buffer); ok {
		t.Errorf("ValidateOligoText returned 'ok' on bad input:\n%s", text)
	}
}

func TestCountLinkers(t *testing.T) {
	text := getFakeOligoText()
	buffer := bufio.NewReader(strings.NewReader(text))
	if numLinkers := CountLinkers(buffer); numLinkers != 1 {
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
