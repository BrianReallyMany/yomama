package sequtil

import (
    "bufio"
    "bytes"
    "io"
)

type QualSeq struct {
    Header string // The sequence header
    Qual   string // The quality string
}

type QualReader struct {
    reader *bufio.Reader

    next  [2][]byte
    index int
}

func NewQualReader(reader *bufio.Reader) *QualReader {
    s := &QualReader{}
    s.reader = reader

    return s
}

// Returns whether or not there is another sequence to be grabbed. Used for iteration.
func (s *QualReader) HasNext() bool {
    var err error

    // Grab next two lines, verify, and return false if any errors occur.

    s.next[0], err = s.reader.ReadBytes('\n')
    if err != nil || len(s.next[0]) == 0 || s.next[0][0] != '>' {
        return false
    }

    s.next[1], err = s.reader.ReadBytes('\n')
    if (err != nil && err != io.EOF) || len(s.next[1]) == 0 {
        return false
    }

    // Sanitize newlines
    s.next[0] = bytes.Trim(s.next[0], "\n")
    s.next[1] = bytes.Trim(s.next[1], "\n")

    return true
}

// Gets the next index and sequence when iterating
func (s *QualReader) Next() (int, QualSeq) {
    defer func(){s.index+=1}()

    return s.index, QualSeq{string(s.next[0][1:]), string(s.next[1])}
}
