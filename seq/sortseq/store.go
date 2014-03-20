package sortseq

import (
    "bytes"
    "encoding/gob"
    . "github.com/BrianReallyMany/yomama/seq"
    "github.com/BrianReallyMany/yomama/seq/sequtil"
    "math"
    "os"
)

// Fixed line width in seq store file.
const storeLineWidth = 120

type storeEntry struct {
    startLine   uint
    headerLines uint
    basesLines  uint
    scoresLines   uint
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
    defer file.Close()
    if err != nil {
        return nil, err
    }

    s := &Store{}

    s.fileName = fileName
    s.seqIndex = make(map[SortKey][]storeEntry)

    // Read the seqIndex from our map file
    mapFile, err := os.Open("."+s.fileName+".map")
    if err == nil {
        decoder := gob.NewDecoder(mapFile)
        decoder.Decode(&s.seqIndex)
    }

    return s, nil
}

func (s *Store) AddSeq(seq Seq) error {
    // Open and verify the file
    file, err := os.OpenFile(s.fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    defer file.Close()
    if err != nil {
        return err
    }

    // Make sure the key exists, make its section if it doesn't
    mySortKey := SortKey{seq.Locus, seq.Sample}
    _, keyExists := s.seqIndex[mySortKey]
    if !keyExists {
        s.seqIndex[mySortKey] = make([]storeEntry, 0, 1)
    }

    // Split data into fixed-width strings
    headerLines := splitFixedWidth([]byte(seq.Header), storeLineWidth)
    basesLines := splitFixedWidth([]byte(seq.Bases), storeLineWidth)
    // Convert Scores from []int to string
    scorestring := seq.ScoresToString()
    scoresLines := splitFixedWidth([]byte(scorestring), storeLineWidth)

    // Add the store entry
    s.seqIndex[mySortKey] = append(s.seqIndex[mySortKey], storeEntry{s.lineCount, uint(len(headerLines)), uint(len(basesLines)), uint(len(scoresLines))})

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

    for _, line := range scoresLines {
        file.Write(line)
        file.WriteString("\n")
    }

    // Set new line count
    s.lineCount += uint(len(headerLines)) + uint(len(basesLines)) + uint(len(scoresLines))

    // Write the seqIndex file for persistence
    s.writeMapFile()

    return nil
}

func (s *Store) FetchSeqs(key SortKey) []Seq {
    // Open and verify the file
    file, err := os.OpenFile(s.fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    defer file.Close()
    if err != nil {
        return nil
    }

    // Get the entries for the current key
    entries, ok := s.seqIndex[key]
    if !ok {
        return nil
    }

    seqs := make([]Seq, len(entries))

    for i, entry := range entries {
        header := make([]byte, entry.headerLines*storeLineWidth)
        bases := make([]byte, entry.basesLines*storeLineWidth)
        scores := make([]byte, entry.scoresLines*storeLineWidth)

        file.Seek(int64(entry.startLine*storeLineWidth), os.SEEK_SET)

        // Grab the header
        for j := 0; j < int(entry.headerLines); j++ {
            file.Read(header[j*storeLineWidth:(j+1)*storeLineWidth]) // Read current line
            file.Seek(1, os.SEEK_CUR) // Skip newline
        }

        // Grab the bases
        for j := 0; j < int(entry.basesLines); j++ {
            file.Read(bases[j*storeLineWidth:(j+1)*storeLineWidth]) // Read current line
            file.Seek(1, os.SEEK_CUR) // Skip newline
        }

        // Grab the scores
        for j := 0; j < int(entry.scoresLines); j++ {
            file.Read(scores[j*storeLineWidth:(j+1)*storeLineWidth]) // Read current line
            file.Seek(1, os.SEEK_CUR) // Skip newline
        }

        // Construct the seq
        seqs[i].Header = string(bytes.Trim(header, " "))
        seqs[i].Bases = string(bytes.Trim(bases, " "))
	// scores is a []byte; Seq.Scores is a []int
        seqs[i].Scores = sequtil.QualStringToIntSlice(string(scores))
        seqs[i].Sample = key.Sample
        seqs[i].Locus = key.Locus
        seqs[i].Reverse = false
    }

    return seqs
}

// writeMapFile writes the store's map into a binary file so it can reload it's seq index map later
func (s *Store) writeMapFile() error {
    file, err := os.Create("."+s.fileName+".map")
    if err != nil {
        return err
    }
    defer file.Close()

    // Create a binary gob encoder to serialize the map
    encoder := gob.NewEncoder(file)

    // Write the map to the file
    if err := encoder.Encode(s.seqIndex); err != nil {
        return err
    }

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
