package sortseq

// SortedSeq is the type that represents a genetic sequence sorted by primer, barcode, and linker.
type SortedSeq struct {
    Header string
    Seq    string
    Qual   string

    Primers  [2]string
    Barcodes [2]string
    Linkers  [2]string
}
