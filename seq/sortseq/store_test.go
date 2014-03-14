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

    testSeq1 := Seq{strings.Repeat("testSeq1", 100), "ATGC", "@#$%", "sample1", "locus1", false}

    // Add the seqs
    store.AddSeq(testSeq1)

    // Fetch the seqs
    seqs := store.FetchSeqs(SortKey{"sample1", "locus1"})

    if len(seqs) != 1 {
        t.Fail()
    }

    if seqs[0] != testSeq1 {
        t.Fail()
    }

    // Cleanup after ourselves
    os.Remove("teststore")
    os.Remove(".teststore.map")
}
