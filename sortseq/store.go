package sortseq

import (
    "errors"
    "math"
    "os"
)

// Fixed line width in seq store file.
const StoreLineWidth = 120

type storeEntry struct {
    startLine   uint
    headerLines uint
    basesLines    uint
    qualLines   uint
}

// Store facilitates the storage and retrieval of sorted seqs
type Store struct {
    fileName string

    seqIndex  map[SortKey]storeEntry // Map sort keys to store entries
    lineCount uint                   // Line count of the storage file
}

func NewStore(fileName string) (*Store, error) {
    // Make sure the provided file exists. Create it if it doesn't.
    _, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        return nil, err
    }

    s := &Store{}

    s.fileName = fileName
    s.seqIndex = make(map[SortKey]storeEntry)

    return s, nil
}

func (s *Store) AddSeq(seq SortedSeq) error {
    // Open and verify the file
    file, err := os.OpenFile(s.fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        return err
    }

    // Make sure the key is unique
    _, keyExists := s.seqIndex[seq.Key]
    if keyExists {
        return errors.New("Seq with key already exists")
    }

    // Split data into fixed-width strings
    headerLines := splitFixedWidth([]byte(seq.Header), StoreLineWidth)
    basesLines := splitFixedWidth([]byte(seq.Bases), StoreLineWidth)
    qualLines := splitFixedWidth([]byte(seq.Qual), StoreLineWidth)

    // Create the store entry
    s.seqIndex[seq.Key] = storeEntry{s.lineCount, uint(len(headerLines)), uint(len(basesLines)), uint(len(qualLines))}

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
	lines := len(str)/fixedWidth + 1
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
