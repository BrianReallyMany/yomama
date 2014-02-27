package sequtil

import (
	"fmt"
    	"strings"
    	"io/ioutil"
)

func ReadOligoFile(filename string) string {
    text, err := ioutil.ReadFile(filename)
    if err != nil {
    	fmt.Println("An error occurred while trying to read oligo file %s", filename)
        fmt.Println(err)
        return ""
    }
    return string(text)
}

func OligoTextToMap(input string) map[[2]string]string {
	lines := strings.Split(input, "\n")
	m := make(map[[2]string]string)
	for _, line := range lines {
		if line != "" {
			var value string;
			fields := strings.Split(line, "\t")
			if len(fields) < 2 {
				continue
			}
			key := [2]string{fields[0], fields[1]}
			if len(fields) == 3 {
				value = fields[2]
			}
			m[key] = value
		}
	}
	return m
}


