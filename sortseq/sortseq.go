package sortseq

import (
    	"strings"
	"github.com/BrianReallyMany/yomama/sortseq/oligo"
)

type SeqSorter struct {
	primerMap map[[2]string]string
	barcodeMap map[[2]string]string
	linkers [][2]string
}

func NewSeqSorter(input string) (*SeqSorter, error) {
	lines := strings.Split(input, "\n")
	primerMap := make(map[[2]string]string)
	barcodeMap := make(map[[2]string]string)
	ok := oligo.ValidateOligoText(input)
	if !ok {
		return &SeqSorter{}, &oligo.OligoError{"Failed to validate oligo file\n"}
	}
	numLinkers := oligo.CountLinkers(input)
	linkers := make([][2]string, numLinkers)
	linkerCount := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		oligotype := oligo.OligoType(line)
		fields := strings.Split(line, "\t")
		switch oligotype {
		case "":
			continue
		case "barcode":
			oligoSeqs := [2]string{fields[1], fields[2]}
			oligoID := fields[3]
			barcodeMap[oligoSeqs] = oligoID
		case "primer":
			oligoSeqs := [2]string{fields[1], fields[2]}
			oligoID := fields[3]
			primerMap[oligoSeqs] = oligoID
		case "linker":
			linkers[linkerCount] = [2]string{fields[1], fields[2]}
			linkerCount++
		}
	}
	sorter := &SeqSorter{primerMap, barcodeMap, linkers}
	return sorter, nil
}

