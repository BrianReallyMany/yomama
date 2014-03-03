package sequtil

import (
    "bufio"
    "strings"
    "testing"
)

func TestQualReaderIteration(t *testing.T) {
    reader := bufio.NewReader(strings.NewReader(">seq1\nA,#%[\n>seq2\n@!)BB\n>seq3\nAAAAA"))

    qualReader := NewQualReader(reader)

    testSeqs := [...] struct {
                          seq QualSeq
                          pass bool
                      } {
        {QualSeq{"seq1", "A,#%["}, false},
        {QualSeq{"seq2", "@!)BB"}, false},
        {QualSeq{"seq3", "AAAAA"}, false},
    }

    for qualReader.HasNext() {
        i, seq := qualReader.Next()

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
