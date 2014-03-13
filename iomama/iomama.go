package iomama

import (
    "bufio"
//    "bytes"
//    "io"
    "github.com/BrianReallyMany/yomama/iomama/fasta"
    "github.com/BrianReallyMany/yomama/iomama/qual"
    . "github.com/BrianReallyMany/yomama/seq"
)


type FastaQualReader struct {
    freader *fasta.FastaReader
    qreader *qual.QualReader

    fseq fasta.FastaSeq
    qseq qual.QualSeq
    next Seq  // Stores next Seq. Used for iteration
}

func NewFastaQualReader(fbuffer, qbuffer *bufio.Reader) *FastaQualReader {
    fq := &FastaQualReader{}
    fq.freader = fasta.NewFastaReader(fbuffer)
    fq.qreader = qual.NewQualReader(qbuffer)

    return fq
}

func (fq *FastaQualReader) HasNext() bool {

    // Grab next seq, verify, and return false if any errors occur.
    if !fq.freader.HasNext() || !fq.qreader.HasNext() {
	    return false
    }

    _, fq.fseq = fq.freader.Next()
    _, fq.qseq = fq.qreader.Next()

    if fq.fseq.Header != fq.qseq.Header {
	    return false
    }
    return true
}

func (fq *FastaQualReader) Next() Seq {
	return Seq{fq.fseq.Header, fq.fseq.Seq, fq.qseq.Qual, "", "", false}
}

