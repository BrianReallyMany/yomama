package iomama

import (
    "strings"
    "testing"
    . "github.com/BrianReallyMany/yomama/seq"
)

func getFastaQualReader() *FastaQualReader {
    fbuffer := strings.NewReader(">seq1\nATGCT\n>seq2\nATGCG\n>seq3\nATGCA")
    qbuffer := strings.NewReader(">seq1\n20 30 30 30 20\n>seq2\n30 35 35 35 40\n>seq3\n40 40 40 40 20")
    fqreader := NewFastaQualReader(fbuffer, qbuffer)
    return fqreader
}


func TestReadFastaQual(t *testing.T) {
	fqreader := getFastaQualReader()
	
	testSeqs := [...] struct {
				seq Seq
				pass bool
			} {
		{Seq{"seq1", "ATGCT", "20 30 30 30 20", "", "", false}, false},
		{Seq{"seq2", "ATGCG", "30 35 35 35 40", "", "", false}, false},
		{Seq{"seq3", "ATGCA", "40 40 40 40 20", "", "", false}, false},
	}
	i := 0
	for fqreader.HasNext() {
		seq := fqreader.Next()

		if seq.Header != testSeqs[i].seq.Header {
			t.Errorf("FastaQual reader returned seq with header %s; expected %s", seq.Header, testSeqs[i].seq.Header)
		}
		if seq.Bases != testSeqs[i].seq.Bases {
			t.Errorf("FastaQual reader returned seq with bases %s; expected %s", seq.Bases, testSeqs[i].seq.Bases)
		}
		if seq.Scores != testSeqs[i].seq.Scores {
			t.Errorf("FastaQual reader returned seq with scores %s; expected %s", seq.Scores, testSeqs[i].seq.Scores)
		}
		i++
	}
}
