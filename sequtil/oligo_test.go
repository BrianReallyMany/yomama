package sequtil

import "testing"

func getFakeOligoText() string {
	result := "GATT\tACA\tlocus1\n"
	result += "GGGG\tCCCC\tlocus2\n"
	result += "ACGT\tTGCA\tsample1\n"
	result += "AAA\tTTT\n"
	return result
}

// TODO define oligo input file spec. Will have to be more complicated
// than this. probably "oligo_type\toligo1\toligo2(optional)\tid\n" kine
// with possibility of only forward linkers, only forward barcodes,
// only forward primers -- and can have "only forward primers; forward
// and reverse barcodes" etc., but can't have "some forward primers; some
// forward and reverse primers". So reading the oligo file
// tells us whether to try to trim only front or front and back.
// but ... if we just try to trim front and back, and possibly use an
// empty string ... no, won't work the way it's written now. hrmph...
func TestOligoTextToMap(t *testing.T) {
	text := getFakeOligoText()
	resultmap := OligoTextToMap(text)
	mapinput := [2]string{"GGGG", "CCCC"}
	mapoutput := resultmap[mapinput]
	if (mapoutput != "locus2") {
		t.Errorf("OligoTextToMap returned a map with m[%s] = %s, wanted 'locus2'", mapinput, mapoutput)
	}
}

func TestOligoTextToMapLinkerEntry(t *testing.T) {
	text := getFakeOligoText()
	resultmap := OligoTextToMap(text)
	mapinput := [2]string{"AAA", "TTT"}
	mapoutput := resultmap[mapinput]
	if (mapoutput != "") {
		t.Errorf("OligoTextToMap returned a map with m[%s] = %s, wanted empty string", mapinput, mapoutput)
	}
}

