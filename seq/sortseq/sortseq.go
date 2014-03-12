package sortseq

import (
	"fmt"
    	"strings"
	. "github.com/BrianReallyMany/yomama/seq"
	"github.com/BrianReallyMany/yomama/seq/sequtil"
	"github.com/BrianReallyMany/yomama/seq/oligo"
)

type SeqSorter struct {
	primerMap map[[2]string]string
	barcodeMap map[[2]string]string
	linkers [][2]string
	SeqSorterOptions
}

type SeqSorterOptions struct {
	bdiffs int
	ldiffs int
	pdiffs int
	checkReverse bool
}

type SeqSorterError struct {
	Problem string
}

func (e *SeqSorterError) Error() string {
	return fmt.Sprintf("SeqSorter error occurred: %s", e.Problem)
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
			linkers[len(linkers)-1] = [2]string{fields[1], fields[2]}
		}
	}
	sorter := &SeqSorter{primerMap, barcodeMap, linkers, SeqSorterOptions{}}
	return sorter, nil
}

func (s *SeqSorter) SortSeq(seq Seq) (Seq, error) {
	// TODO be flexible; if seq already has sample, skip debarcoding...
	// Find barcode pair with best match
	barcodeKeys := getSliceOfKeys(s.barcodeMap)
	bestBarcodes, mismatches := bestMatch(barcodeKeys, seq.Bases)
	if mismatches > s.bdiffs {
		return Seq{}, &SeqSorterError{"Exceeded maximum number of differences between barcode and sequence"}
	}
	sampleName, ok  := s.barcodeMap[bestBarcodes]
	if !ok {
		return Seq{}, &SeqSorterError{"Couldn't find sample name to match barcode pair. This is strange."}
	}





	// trim barcodes off seq bases and qual scores

	// TODO flag for "check reversed linker/primer pairs"?
	// Find linker pair with best match
	// trim linkers off seq bases and qual scores

	// Find primer pair with best match
	// get sample name
	// trim primers off seq bases and qual scores

	// Make new Seq (or update Seq) with all this info
	// Return it.
	// TODO these are here so it builds...
	fmt.Println(sampleName)
	fmt.Println(mismatches)
	return Seq{"", "gattaca", "", "", "", true}, nil
}

// From a list of barcode pairs, return the pair that matches
// most closely, and the number of bases that don't match
func bestMatch(oligos [][2]string, seq string) ([2]string, int) {
	winner := [2]string{"", ""}
	mismatches := len(seq) + 1
	for _, pair := range oligos {
		if misses := sequtil.MatchBeginAndEnd(pair, seq); misses < mismatches {
			mismatches = misses
			winner = pair
		}
	}
	return winner, mismatches
}

func getSliceOfKeys(m map[[2]string]string) [][2]string {
	mySlice := make([][2]string, len(m))
	i := 0
	for k := range m {
		mySlice[i] = k
		i++
	}
	return mySlice
}
