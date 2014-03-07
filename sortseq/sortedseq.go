package sortseq

type SortKey struct {
    Primer  string
    Barcode string
    Linker  string
}

// SortedSeq is the type that represents a genetic sequence sorted by primer, barcode, and linker.
type SortedSeq struct {
    Header string
    Bases  string
    Qual   string
    Key    SortKey
}
