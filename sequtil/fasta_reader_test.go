package sequtil

import (
    "bufio"
    "strings"
    "testing"
)

func TestFastaReaderIteration(t *testing.T) {
    reader := bufio.NewReader(strings.NewReader(">seq1\nATGCT\n>seq2\nATGCG\n>seq3\nATGCA"))

    fastaReader := NewFastaReader(reader)

    testSeqs := [...] struct {
                          seq FastaSeq
                          pass bool
                      } {
        {FastaSeq{"seq1", "ATGCT"}, false},
        {FastaSeq{"seq2", "ATGCG"}, false},
        {FastaSeq{"seq3", "ATGCA"}, false},
    }

    for fastaReader.HasNext() {
        i, seq := fastaReader.Next()

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
