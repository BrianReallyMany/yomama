package sequtil

import (
    "bufio"
    "bytes"
    "io"
)

type Seq struct {
    Header string
    Seq string
    Scores string
}

type SeqReader struct {
    fastaReader *bufio.Reader

    next [2][]byte // Stores next 2 lines. Used for iteration
    index int // Stores next index when iterating
}

func NewSeqReader(fastaReader *bufio.Reader) *SeqReader {
    s := &SeqReader{}
    s.fastaReader = fastaReader
    
    return s
}

// Returns whether or not there is another sequence to be grabbed. Used for iteration.
func (s *SeqReader) HasNext() bool {
    var err error

    // Grab next two lines, verify, and return false if any errors occur.

    s.next[0], err = s.fastaReader.ReadBytes('\n')
    if err != nil || len(s.next[0]) == 0 || s.next[0][0] != '>' {
        return false
    }

    s.next[1], err = s.fastaReader.ReadBytes('\n')
    if (err != nil && err != io.EOF) || len(s.next[1]) == 0 {
        return false
    }

    // Sanitize newlines
    s.next[0] = bytes.Trim(s.next[0], "\n")
    s.next[1] = bytes.Trim(s.next[1], "\n")

    return true
}

// Gets the next index and sequence when iterating
func (s *SeqReader) Next() (int, Seq) {
    defer func(){s.index+=1}()

    return s.index, Seq{string(s.next[0][1:]), string(s.next[1]), ""}
}

