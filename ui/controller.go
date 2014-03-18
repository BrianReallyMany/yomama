package ui

import (
	"bufio"
	"strconv"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"github.com/BrianReallyMany/yomama/dozens"
	"github.com/BrianReallyMany/yomama/iomama"
	"github.com/BrianReallyMany/yomama/iomama/fastaqual"
	"github.com/BrianReallyMany/yomama/iomama/fastq"
	"github.com/BrianReallyMany/yomama/seq/sortseq"
)

type MamaController struct {
}

// Instantiate a new yomama controller
func MakeMamaController() *MamaController {
	return &MamaController{}
}

func (c *MamaController) Dozens() string {
	return dozens.RandomDozens()
}

func (c *MamaController) PrepFiles(args []string, ch chan string) {
	if len(args) < 1 {
		return
	}

	myPath := args[0]

	ch <- "\nVerifying files...\n"

	// Verify folder exists, contains .fasta, .qual and .oligo files
	var fastqFileName, fastaFileName, qualFileName, oligoFileName string

	fastqfiles, err := filepath.Glob(myPath + "/*.fastq")
	if err != nil || len(fastqfiles) != 1 {
		ch <- "No fastq files found, checking for fasta and qual files...\n"

		fastafiles, err := filepath.Glob(myPath + "/*.fasta")
		if err != nil || len(fastafiles) != 1 {
			ch <- "PrepFiles: locating fasta file failed"
			ch <- err.Error()
			return
		}
		qualfiles, err := filepath.Glob(myPath + "/*.qual")
		if err != nil || len(qualfiles) != 1 {
			ch <- "PrepFiles: locating qual file failed"
			ch <- err.Error()
			return
		}

		fastaFileName = fastafiles[0]
		qualFileName = qualfiles[0]
	} else {
		fastqFileName = fastqfiles[0]
	}

	oligofiles, err := filepath.Glob(myPath + "/*.oligo")
	if err != nil || len(oligofiles) != 1 {
		ch <- "PrepFiles: locating oligo file failed"
		ch <- err.Error()
		return
	} else {
		oligoFileName = oligofiles[0]
	}

	var seqReader iomama.SeqReader

	if fastqFileName != "" {
		fastqfile, err := os.Open(fastqFileName)
		if err != nil {
			ch <- "PrepFiles: fastq open failed"
			ch <- err.Error()
			return
		}

		seqReader = fastq.NewFastqReader(bufio.NewReader(fastqfile))
	} else if fastaFileName != "" && qualFileName != "" {
		fastafile, err := os.Open(fastaFileName)
		if err != nil {
			ch <- "PrepFiles: fasta open failed"
			ch <- err.Error()
			return
		}
		qualfile, err := os.Open(qualFileName)
		if err != nil {
			return
		}

		// Make FastaQualReader, SeqSorter
		seqReader = fastaqual.NewFastaQualReader(fastafile, qualfile)
	}

	if err != nil {
		return
	}

	oligofile, err := os.Open(oligoFileName)
	temporaryDefaultSeqSorterOptions := sortseq.NewSeqSorterOptions(0, 0, 0, true)
	sorter, err := sortseq.NewSeqSorter(oligofile, temporaryDefaultSeqSorterOptions)
	if err != nil {
		return
	}

	// Make Store
	store, err := sortseq.NewStore(myPath + "/yomama.store")
	if err != nil {
		return
	}

	// Track where sorting fails and why
	errorsMap := make(map[string]int)

	ch <- "Reading in sequences..."

	for seqReader.HasNext() {
		seq := seqReader.Next()
		sortedseq, err := sorter.SortSeq(seq)
		if err != nil {
			// Add entry to errorsMap
			// (required type assertion)
			if e, ok := err.(*sortseq.SeqSorterError); ok {
				errorsMap[e.Where] += 1
			}
			continue
		}
		err = store.AddSeq(sortedseq)
		if err != nil {
			ch <- "PrepFiles: error storing seq--"
			ch <- sortedseq.ToString()
		}
	}
	ch <- "...done.\n"
	ch <- "Error summary:"
	ch <- "Type of error\tCount"
	for k, v := range errorsMap {
		ch <- k + "\t" + strconv.Itoa(v) + "\n"
	}
	ch <- "END"
}

func (c *MamaController) System(args []string) string {
	out, err := exec.Command(args[0], args[1:]...).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}
