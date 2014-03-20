package fastq

import (
    "bufio"
    . "github.com/BrianReallyMany/yomama/seq"
    "github.com/BrianReallyMany/yomama/seq/sequtil"
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
        {Seq{"seq1", "ATGCT", sequtil.StringToPhredScoreSlice("#$%^&", false), "", "", false}, false},
        {Seq{"seq2", "ATGCG", sequtil.StringToPhredScoreSlice("&^%$#", false), "", "", false}, false},
        {Seq{"seq3", "ATGCA", sequtil.StringToPhredScoreSlice("ABCDE", false), "", "", false}, false},
    }

    i := 0
    for fastqReader.HasNext() {
        seq := fastqReader.Next()

        if seq.Bases == testSeqs[i].seq.Bases && seq.Header == testSeqs[i].seq.Header {
            testSeqs[i].pass = true
        }
	for j, score := range seq.Scores {
		if score != testSeqs[i].seq.Scores[j] {
			testSeqs[i].pass = false
		}
	}
        i++
    }

    for _, seq := range testSeqs {
        if !seq.pass {
            t.FailNow()
        }
    }
}
