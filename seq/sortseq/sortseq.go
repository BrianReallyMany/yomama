package sortseq

import (
    	"strings"
	. "github.com/BrianReallyMany/yomama/seq"
	"github.com/BrianReallyMany/yomama/seq/sequtil"
	"github.com/BrianReallyMany/yomama/seq/oligo"
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
	sorter := &SeqSorter{primerMap, barcodeMap, linkers}
	return sorter, nil
}

func (s *SeqSorter) SortSeq(seq Seq) SortedSeq {
	// Find barcode pair with best match
	// get sample name
	// trim barcodes off seq bases and qual scores

	// Find linker pair with best match
	// trim linkers off seq bases and qual scores

	// Find primer pair with best match
	// get sample name
	// trim primers off seq bases and qual scores

	// Make SortKey
	// Make SortedSeq
	// Return it.
	seq.Bases = "gattaca"
	return SortedSeq{seq, SortKey{}}
}

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
