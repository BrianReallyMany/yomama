package sortseq

import (
	"io"
	"bufio"
	"fmt"
	"errors"
    	"strings"
	"strconv"
	. "github.com/BrianReallyMany/yomama/seq"
	"github.com/BrianReallyMany/yomama/seq/sequtil"
)

type SeqSorter struct {
	primerMap map[[2]string]string
	barcodeMap map[[2]string]string
	linkers [][2]string
	*SeqSorterOptions
}

type SeqSorterOptions struct {
	bdiffs int
	ldiffs int
	pdiffs int
	checkReverse bool
}

type SeqSorterError struct {
	Problem string
	Where string
}

func (e *SeqSorterError) Error() string {
	return fmt.Sprintf("SeqSorter error occurred: %s", e.Problem)
}

func NewSeqSorter(reader io.Reader, opts *SeqSorterOptions) (*SeqSorter, error) {
	primerMap := make(map[[2]string]string)
	barcodeMap := make(map[[2]string]string)
	linkers := make([][2]string, 0)
	checkrev := opts.checkReverse

	// Process oligo file line by line
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		fields, err := validateOligoLine(line)
		if err != nil {
			return &SeqSorter{}, err
			// TODO better error
		}

		// Don't fail just because you found an empty line
		if len(fields) == 0 {
			continue
		}

		oligotype := fields[0]
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
			if checkrev {
				oligoSeqs = [2]string{fields[2], fields[1]}
				primerMap[oligoSeqs] = oligoID
			}
		case "linker":
			linkers = append(linkers, [2]string{fields[1], fields[2]})
			if checkrev {
				linkers = append(linkers, [2]string{fields[2], fields[1]})
			}
		}
	}
	sorter := &SeqSorter{primerMap, barcodeMap, linkers, opts}
	return sorter, nil
}

func validateOligoLine(line string) ([]string, error) {
	line = strings.Trim(line, "\n ")
	fields := strings.Split(line, "\t")
	switch oligotype := fields[0]; oligotype {
	case "barcode":
		if (len(fields) != 4 || fields[1] == "" || fields[3] == "") {
			return fields, errors.New("bad barcode")
			// TODO better errors
		}
		return fields, nil
	case "primer":
		if (len(fields) != 4 || fields[1] == "" || fields[3] == "") {
			return fields, errors.New("bad primer")
		}
		return fields, nil
	case "linker":
		if len(fields) < 2 {
			return fields, errors.New("bad linker")
		}
		return fields, nil
	default:
		fmt.Printf("Unkown oligo type encountered: %s\n", oligotype)
	}
	return fields, errors.New("Invalid oligo line.")
}

func oligoType(line string) string {
	fields := strings.Split(line, "\t")
	return fields[0]
}

func (s *SeqSorter) ToString() string {
	result := "primerMap: " + strconv.Itoa(len(s.primerMap)) + " entries\n"
	result += "barcodeMap: " + strconv.Itoa(len(s.barcodeMap)) + " entries\n"
	result += "linkers: " + strconv.Itoa(len(s.linkers)) + " entries\n"
	result += "options:\n"
	result += "\tbdiffs=" + strconv.Itoa(s.bdiffs) + "\n"
	result += "\tldiffs=" + strconv.Itoa(s.ldiffs) + "\n"
	result += "\tpdiffs=" + strconv.Itoa(s.pdiffs) + "\n"
	result += "\tcheckReverse=" + strconv.FormatBool(s.checkReverse)
	return result
}

func (s *SeqSorter) SortSeq(seq Seq) (Seq, error) {
	// TODO be flexible; if seq already has sample, skip debarcoding...
	// Find barcode pair with best match
	barcodeKeys := getSliceOfKeys(s.barcodeMap)
	bestBarcodes, mismatches := bestMatch(barcodeKeys, seq.Bases)
	// Verify acceptable number of mismatches
	if mismatches > s.bdiffs {
		return seq, &SeqSorterError{"Exceeded maximum number of differences between barcode and sequence", "barcode"}
	}
	// Get sample name
	sampleName, ok  := s.barcodeMap[bestBarcodes]
	if !ok {
		return seq, &SeqSorterError{"Couldn't find sample name to match barcode pair. This is strange.", "barcode"}
	}
	// Trim barcodes off seq bases and qual scores
	err := seq.TrimEnds(len(bestBarcodes[0]), len(bestBarcodes[1]))
	if err != nil {
		return seq, &SeqSorterError{"Error trimming ends of seq.", "barcode"}
	}

	// TODO flag for "check reversed linker/primer pairs"?

	// Find linker pair with best match
	bestLinkers, mismatches := bestMatch(s.linkers, seq.Bases)
	if mismatches > s.ldiffs {
		return seq, &SeqSorterError{"Exceeded maximum number of differences between linker and sequence", "linker"}
	}
	// trim linkers off seq bases and qual scores
	err = seq.TrimEnds(len(bestLinkers[0]), len(bestLinkers[1]))
	if err != nil {
		return seq, &SeqSorterError{"Error trimming ends of debarcoded seq.", "linker"}
	}

	// Find primer pair with best match
	primerKeys := getSliceOfKeys(s.primerMap)
	bestPrimers, mismatches := bestMatch(primerKeys, seq.Bases)
	if mismatches > s.pdiffs {
		return seq, &SeqSorterError{"Exceeded maximum number of differences between primer and sequence", "primer"}
	}
	locus, ok := s.primerMap[bestPrimers]
	if !ok {
		return seq, &SeqSorterError{"Couldn't find locus name to match primer pair.", "primer"}
	}
	// Trim primers off seq bases and qual scores
	err = seq.TrimEnds(len(bestPrimers[0]), len(bestPrimers[1]))
	if err != nil {
		return seq, &SeqSorterError{"Error trimming ends of delinkered seq.", "primer"}
	}

	// Update seq info and return
	seq.Sample = sampleName
	seq.Locus = locus
	return seq, nil
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
