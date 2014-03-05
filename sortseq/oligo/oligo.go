package oligo

import (
	"fmt"
    	"strings"
    	"io/ioutil"
)

type SeqSorter struct {
	primerMap map[[2]string]string
	barcodeMap map[[2]string]string
	linkers [][2]string
}


func ReadOligoFile(filename string) string {
    text, err := ioutil.ReadFile(filename)
    if err != nil {
    	fmt.Printf("An error occurred while trying to read oligo file %s\n", filename)
        fmt.Println(err)
        return ""
    }
    return string(text)
}

func OligoTextToSeqSorter(input string) (SeqSorter, error) {
	lines := strings.Split(input, "\n")
	primerMap := make(map[[2]string]string)
	barcodeMap := make(map[[2]string]string)
	ok := ValidateOligosText(input)
	if !ok {
		fmt.Println("Failed to validate oligo file\n")
		return SeqSorter{}, nil
	}
	numLinkers := CountLinkers(input)
	linkers := make([][2]string, numLinkers)
	linkerCount := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		oligotype := OligoType(line)
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
	sorter := SeqSorter{primerMap, barcodeMap, linkers}
	return sorter, nil
}

// Returns number of linker lines and whether the file is valid
func ValidateOligosText(input string) bool {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		if line != "" {
			ok := ValidateOligoLine(line)
			if !ok {
				fmt.Printf("Oligo file is invalid -- problem at line %d\n", i)
				return false
			}
		}
	}
	return true
}

func CountLinkers(input string) int {
	lines := strings.Split(input, "\n")
	numLinkers := 0
	for _, line := range lines {
		if line != "" {
			oligoType := OligoType(line)
			if oligoType == "linker" {
				numLinkers++
			}
		}
	}
	return numLinkers
}

// Returns oligo type if line is valid, empty string if not
func ValidateOligoLine(line string) bool {
	fields := strings.Split(line, "\t")
	switch oligotype := fields[0]; oligotype {
	case "barcode":
		if (len(fields) != 4 || fields[1] == "" || fields[3] == "") {
			return false
		}
		return true
	case "primer":
		if (len(fields) != 4 || fields[1] == "" || fields[3] == "") {
			return false
		}
		return true
	case "linker":
		if len(fields) < 2 {
			return false
		}
		return true
	default:
		fmt.Printf("Unkown oligo type encountered: %s\n", oligotype)
	}
	return false
}

func OligoType(line string) string {
	fields := strings.Split(line, "\t")
	return fields[0]
}
