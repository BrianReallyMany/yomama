package sequtil

import (
    "bufio"
)

type QualSeq struct {
    Header string // The sequence header
    Qual   string // The quality string
}

type QualReader struct {
    reader bufio.Reader

    index int
    next [2][]byte
}
