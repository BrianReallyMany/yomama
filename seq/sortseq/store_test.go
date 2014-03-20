package sortseq

import (
    . "github.com/BrianReallyMany/yomama/seq"
    "os"
    "strings"
    "testing"
)

func TestStore_AddFetchSeq(t *testing.T) {
    store, err := NewStore("teststore")
    if err != nil {
        t.Fail()
    }

    testSeq1 := Seq{strings.Repeat("testSeq1", 100), "ATGC", []int{25, 30, 30, 28}, "sample1", "locus1", false}

    // Add the seqs
    store.AddSeq(testSeq1)

    // Fetch the seqs
    seqs := store.FetchSeqs(SortKey{"sample1", "locus1"})

    if len(seqs) != 1 {
        t.Fail()
    }

    result := seqs[0]
    if result.Bases != testSeq1.Bases {
        t.Fail()
    }
    for i, score := range testSeq1.Scores {
	    if score != result.Scores[i] {
		    t.Fail()
	    }
    }
    if result.Locus != testSeq1.Locus {
	    t.Fail()
    }
    if result.Sample != testSeq1.Sample {
	    t.Fail()
    }

    // Cleanup after ourselves
    os.Remove("teststore")
    os.Remove(".teststore.map")
}
