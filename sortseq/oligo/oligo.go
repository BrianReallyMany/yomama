package oligo

import (
	"fmt"
    	"strings"
    	"io/ioutil"
)

func ReadOligoFile(filename string) string {
    text, err := ioutil.ReadFile(filename)
    if err != nil {
    	fmt.Printf("An error occurred while trying to read oligo file %s\n", filename)
        fmt.Println(err)
        return ""
    }
    return string(text)
}

func OligoTextToMapsAndLinkerSlice(input string) (map[[2]string]string, map[[2]string]string, [][2]string) {
	lines := strings.Split(input, "\n")
	primerMap := make(map[[2]string]string)
	barcodeMap := make(map[[2]string]string)
	numLinkers, ok := ValidateOligosText(input)
	if !ok {
		fmt.Println("Failed to validate oligo file\n")
		return nil, nil, nil
	}
	linkers := make([][2]string, numLinkers)
	linkerCount := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		oligotype := ValidateOligoLine(line)
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
	return primerMap, barcodeMap, linkers
}

// Returns number of linker lines and whether the file is valid
func ValidateOligosText(input string) (int, bool) {
	lines := strings.Split(input, "\n")
	numLinkers := 0
	for i, line := range lines {
		if line != "" {
			oligoType := ValidateOligoLine(line)
			if oligoType == "" {
				fmt.Printf("Oligo file is invalid -- problem at line %d\n", i)
				return -1, false
			} else if oligoType == "linker" {
				numLinkers++
			}
		}
	}
	return numLinkers, true
}

// Returns oligo type if line is valid, empty string if not
func ValidateOligoLine(line string) string {
	fields := strings.Split(line, "\t")
	switch oligotype := fields[0]; oligotype {
	case "barcode":
		if (len(fields) != 4 || fields[1] == "" || fields[3] == "") {
			return ""
		}
		return oligotype
	case "primer":
		if (len(fields) != 4 || fields[1] == "" || fields[3] == "") {
			return ""
		}
		return oligotype
	case "linker":
		if len(fields) < 2 {
			return ""
		}
		return oligotype
	default:
		fmt.Printf("Unkown oligo type encountered: %s\n", oligotype)
	}
	return ""
}

