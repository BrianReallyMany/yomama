package sequtil

import (
    "bufio"
    "strings"
    "testing"
)

func TestSeqReaderIteration(t *testing.T) {
    reader := bufio.NewReader(strings.NewReader(">seq1\nATGCT\n>seq2\nATGCG\n>seq3\nATGCA"))

    seqReader := NewSeqReader(reader)

    testSeqs := [...] struct {
                          seq Seq
                          pass bool
                      } {
        {Seq{"seq1", "ATGCT", ""}, false},
        {Seq{"seq2", "ATGCG", ""}, false},
        {Seq{"seq3", "ATGCA", ""}, false},
    }

    for seqReader.HasNext() {
        i, seq := seqReader.Next()

        if seq == testSeqs[i].seq {
            testSeqs[i].pass = true
        }
    }

    for _, seq := range testSeqs {
        if !seq.pass {
            t.FailNow()
        }
    }
}
