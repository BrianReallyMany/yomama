package ui

import (
	"bufio"
	"os"
	"fmt"
	"path/filepath"
    	"log"
    	"os/exec"
	"github.com/BrianReallyMany/yomama/dozens"
	"github.com/BrianReallyMany/yomama/seq/sortseq"
	"github.com/BrianReallyMany/yomama/iomama"
	"github.com/BrianReallyMany/yomama/iomama/fastaqual"
	"github.com/BrianReallyMany/yomama/iomama/fastq"
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

func (c *MamaController) PrepFiles(args []string) {
	if len(args) < 1 {
		return
	}

	myPath := args[0]

	fmt.Println("Verifying files...\n")

	// Verify folder exists, contains .fasta, .qual and .oligo files
	var fastqFileName, fastaFileName, qualFileName, oligoFileName string

	fastqfiles, err := filepath.Glob(myPath + "/*.fastq")
	if err != nil || len(fastqfiles) != 1 {
		fmt.Println("No fastq files found, checking for fasta and qual files...")

		fastafiles, err := filepath.Glob(myPath + "/*.fasta")
		if err != nil || len(fastafiles) != 1 {
			fmt.Println("PrepFiles: locating fasta file failed")
			fmt.Println(err)
			return
		}
		qualfiles, err := filepath.Glob(myPath + "/*.qual")
		if err != nil || len(qualfiles) != 1 {
			fmt.Println("PrepFiles: locating qual file failed")
			fmt.Println(err)
			return
		}

		fastaFileName = fastafiles[0]
		qualFileName = qualfiles[0]
	} else {
		fastqFileName = fastqfiles[0]
	}

	oligofiles, err := filepath.Glob(myPath + "/*.oligo")
	if err != nil || len(oligofiles) != 1 {
		fmt.Println("PrepFiles: locating oligo file failed")
		fmt.Println(err)
		return
	} else {
		oligoFileName = oligofiles[0]
	}

	var seqReader iomama.SeqReader

	if fastqFileName != "" {
		fastqfile, err := os.Open(fastqFileName)
		if err != nil {
			fmt.Println("PrepFiles: fastq open failed")
			fmt.Println(err)
			return
		}

		seqReader = fastq.NewFastqReader(bufio.NewReader(fastqfile))
	} else if fastaFileName != "" && qualFileName != "" {
		fastafile, err := os.Open(fastaFileName)
		if err != nil {
			fmt.Println("PrepFiles: fasta open failed")
			fmt.Println(err)
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
	sorter, err := sortseq.NewSeqSorter(oligofile)
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

	fmt.Println("Reading in sequences...")

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
			fmt.Println("PrepFiles: error storing seq--")
			fmt.Println(sortedseq.ToString())
		}
	}
	fmt.Println("...done.\n")
	fmt.Println("Error summary:")
	fmt.Println("Type of error\tCount")
	for k, v := range errorsMap {
		fmt.Printf("%s\t%d\n", k, v)
	}
}

func (c *MamaController) System(args []string) string {
        out, err := exec.Command(args[0], args[1:]...).Output()
	    if err != nil {
		    log.Fatal(err)
	    }
	return string(out)
}
