package fastq

import (
    "bufio"
    . "github.com/BrianReallyMany/yomama/seq"
    "strings"
    "testing"
)

func TestFastqReaderIteration(t *testing.T) {
    reader := bufio.NewReader(strings.NewReader("@seq1\nATGCT\n+ blah blah  \n#$%^&\n@seq2\nATGCG\n+\n&^%$#\n@seq3\nATGCA\n+  \nABCDE"))

    fastqReader := NewFastqReader(reader)

    testSeqs := [...] struct {
                          seq Seq
                          pass bool
                      } {
        {Seq{"seq1", "ATGCT", "#$%^&", "", "", false}, false},
        {Seq{"seq2", "ATGCG", "&^%$#", "", "", false}, false},
        {Seq{"seq3", "ATGCA", "ABCDE", "", "", false}, false},
    }

    i := 0
    for fastqReader.HasNext() {
        seq := fastqReader.Next()

        if seq == testSeqs[i].seq {
            testSeqs[i].pass = true
        }
        i++
    }

    for _, seq := range testSeqs {
        if !seq.pass {
            t.FailNow()
        }
    }
}
