package fastq

import (
    "bufio"
    "bytes"
    . "github.com/BrianReallyMany/yomama/seq"
    "io"
)

type FastqReader struct {
    reader *bufio.Reader

    next [3][]byte // Stores next 3 meaningful lines. Used for iteration
}

func NewFastqReader(reader *bufio.Reader) *FastqReader {
    f := &FastqReader{}
    f.reader = reader

    return f
}

// Returns whether or not there is another sequence to be grabbed. Used for iteration.
func (f *FastqReader) HasNext() bool {
    var err error

    // Grab next two lines, verify, and return false if any errors occur.

    f.next[0], err = f.reader.ReadBytes('\n')
    if err != nil || len(f.next[0]) == 0 || f.next[0][0] != '@' {
        return false
    }

    f.next[1], err = f.reader.ReadBytes('\n')
    if err != nil {
        return false
    }

    plusLine, err := f.reader.ReadBytes('\n')
    if err != nil || len(plusLine) == 0 || plusLine[0] != '+' {
        return false
    }

    f.next[2], err = f.reader.ReadBytes('\n')
    if (err != nil && err != io.EOF) || len(f.next[2]) == 0 {
        return false
    }

    // Sanitize newlines
    f.next[0] = bytes.Trim(f.next[0], "\n")
    f.next[1] = bytes.Trim(f.next[1], "\n")
    f.next[2] = bytes.Trim(f.next[2], "\n")

    return true
}

// Gets the next index and sequence when iterating
func (f *FastqReader) Next() Seq {
    return Seq{string(f.next[0][1:]), string(f.next[1]), string(f.next[2]), "", "", false}
}
