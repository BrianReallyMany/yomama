package oligo

import (
	"io"
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

func ValidateOligoText(reader io.Reader) bool {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			ok := ValidateOligoLine(line)
			if !ok {
				return false
			}
		}
	}
	return true
}

func CountLinkers(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)
	numLinkers := 0
	for scanner.Scan() {
		line := scanner.Text()
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

