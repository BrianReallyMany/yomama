package sequtil

import (
    "bufio"
    "bytes"
    "io"
)

type FastaSeq struct {
    Header string
    Seq string
}

type FastaReader struct {
    reader *bufio.Reader

    next [2][]byte // Stores next 2 lines. Used for iteration
    index int // Stores next index when iterating
}

func NewFastaReader(reader *bufio.Reader) *FastaReader {
    f := &FastaReader{}
    f.reader = reader

    return f
}

// Returns whether or not there is another sequence to be grabbed. Used for iteration.
func (f *FastaReader) HasNext() bool {
    var err error

    // Grab next two lines, verify, and return false if any errors occur.

    f.next[0], err = f.reader.ReadBytes('\n')
    if err != nil || len(f.next[0]) == 0 || f.next[0][0] != '>' {
        return false
    }

    f.next[1], err = f.reader.ReadBytes('\n')
    if (err != nil && err != io.EOF) || len(f.next[1]) == 0 {
        return false
    }

    // Sanitize newlines
    f.next[0] = bytes.Trim(f.next[0], "\n")
    f.next[1] = bytes.Trim(f.next[1], "\n")

    return true
}

// Gets the next index and sequence when iterating
func (f *FastaReader) Next() (int, FastaSeq) {
    defer func(){f.index+=1}()

    return f.index, FastaSeq{string(f.next[0][1:]), string(f.next[1])}
}
