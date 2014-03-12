package sortseq

import (
    . "github.com/BrianReallyMany/yomama/seq"
    "math"
    "os"
)

// Fixed line width in seq store file.
const StoreLineWidth = 120

type storeEntry struct {
    startLine   uint
    headerLines uint
    basesLines  uint
    qualLines   uint
}

// Store facilitates the storage and retrieval of sorted seqs
type Store struct {
    fileName string

    seqIndex  map[SortKey][]storeEntry // Map sort keys to store entries
    lineCount uint                   // Line count of the storage file
}

func NewStore(fileName string) (*Store, error) {
    // Make sure the provided file exists. Create it if it doesn't.
    file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    s := &Store{}

    s.fileName = fileName
    s.seqIndex = make(map[SortKey][]storeEntry)

    return s, nil
}

func (s *Store) AddSeq(seq Seq) error {
    // Open and verify the file
    file, err := os.OpenFile(s.fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        return err
    }
    defer file.Close()

    // Make sure the key exists, make its section if it doesn't
    mySortKey := SortKey{seq.Locus, seq.Sample}
    _, keyExists := s.seqIndex[mySortKey]
    if !keyExists {
        s.seqIndex[mySortKey] = make([]storeEntry, 0, 1)
    }

    // Split data into fixed-width strings
    headerLines := splitFixedWidth([]byte(seq.Header), StoreLineWidth)
    basesLines := splitFixedWidth([]byte(seq.Bases), StoreLineWidth)
    qualLines := splitFixedWidth([]byte(seq.Scores), StoreLineWidth)

    // Add the store entry
    s.seqIndex[mySortKey] = append(s.seqIndex[mySortKey], storeEntry{s.lineCount, uint(len(headerLines)), uint(len(basesLines)), uint(len(qualLines))})

    // Write to the store file
    file.Seek(0, 2) // Go to the end of the file

    for _, line := range headerLines {
        file.Write(line)
        file.WriteString("\n")
    }

    for _, line := range basesLines {
        file.Write(line)
        file.WriteString("\n")
    }

    for _, line := range qualLines {
        file.Write(line)
        file.WriteString("\n")
    }

    // Set new line count
    s.lineCount += uint(len(headerLines)) + uint(len(basesLines)) + uint(len(qualLines))

    return nil
}

// splitFixedWidth splits a byte slice into multiple slices with a length of the provided fixed
// width. Any extra space on the last line will be filled with whitespaces.
func splitFixedWidth(str []byte, fixedWidth int) [][]byte {
	lines := len(str)/fixedWidth
        if len(str)%fixedWidth != 0 {
            lines++
        }

	splitStr := make([][]byte, lines)
	for i := 0; i < lines; i++ {
		toCopy := []byte(str)[i*fixedWidth : int(math.Min(float64((i+1)*fixedWidth), float64(len(str))))]

		splitStr[i] = make([]byte, fixedWidth)
		copy(splitStr[i], toCopy)
		for j := 0; j < fixedWidth-len(toCopy); j++ {
			splitStr[i][len(toCopy)+j] = ' '
		}
	}

	return splitStr
}
