package oligo

import (
	"fmt"
    	"strings"
    	"io/ioutil"
)

type OligoError struct {
	Problem string
}

func (e *OligoError) Error() string {
	return fmt.Sprintf("oligo problem was %s", e.Problem)
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

func ValidateOligoText(input string) bool {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line != "" {
			ok := ValidateOligoLine(line)
			if !ok {
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

