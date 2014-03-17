package fastaqual

import (
    "bufio"
    "io"
    "github.com/BrianReallyMany/yomama/iomama/fastaqual/fasta"
    "github.com/BrianReallyMany/yomama/iomama/fastaqual/qual"
    . "github.com/BrianReallyMany/yomama/seq"
)

type FastaQualReader struct {
    freader *fasta.FastaReader
    qreader *qual.QualReader

    fseq fasta.FastaSeq
    qseq qual.QualSeq
    next Seq  // Stores next Seq. Used for iteration
}

func NewFastaQualReader(fbuffer, qbuffer io.Reader) *FastaQualReader {
    fq := &FastaQualReader{}
    fq.freader = fasta.NewFastaReader(bufio.NewReader(fbuffer))
    fq.qreader = qual.NewQualReader(bufio.NewReader(qbuffer))

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

