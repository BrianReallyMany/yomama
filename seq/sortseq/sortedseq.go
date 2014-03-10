package sortseq

import (
    . "github.com/BrianReallyMany/yomama/seq"
)

type SortKey struct {
    Locus  string
    Sample string
}

// SortedSeq is the type that represents a genetic sequence sorted by primer, barcode, and linker.
type SortedSeq struct {
    Seq
    SortKey
}
