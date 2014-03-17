package oligo

import (
	"bufio"
	"fmt"
    	"strings"
)

type OligoError struct {
	Problem string
}

func (e *OligoError) Error() string {
	return fmt.Sprintf("oligo problem was %s", e.Problem)
}

func ValidateOligoText(reader *bufio.Reader) bool {
	var err error
	var line []byte
	for ; err == nil; line, err = reader.ReadBytes('\n') {
		sline := string(line)
		if sline != "" {
			ok := ValidateOligoLine(sline)
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
	line = strings.Trim(line, "\n ")
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

