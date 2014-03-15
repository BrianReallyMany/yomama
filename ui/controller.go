package ui

import (
	"os"
	"fmt"
	"path/filepath"
	"io/ioutil"
	"github.com/BrianReallyMany/yomama/dozens"
	"github.com/BrianReallyMany/yomama/seq/sortseq"
	"github.com/BrianReallyMany/yomama/iomama"
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
	oligofiles, err := filepath.Glob(myPath + "/*.oligo")
	if err != nil || len(oligofiles) != 1 {
		fmt.Println("PrepFiles: locating oligo file failed")
		fmt.Println(err)
		return
	}

	fastafile, err := os.Open(fastafiles[0])
	if err != nil {
		fmt.Println("PrepFiles: fasta open failed")
		fmt.Println(err)
		return
	}
	qualfile, err := os.Open(qualfiles[0])
	if err != nil {
		return
	}
	
	// Make FastaQualReader, SeqSorter
	fqreader := iomama.NewFastaQualReader(fastafile, qualfile)

	oligostring, err := ioutil.ReadFile(oligofiles[0])
	if err != nil {
		return
	}
	sorter, err := sortseq.NewSeqSorter(string(oligostring)) // takes string representation of oligo file!
	// TODO consistent io handling for sortseq!
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

	for fqreader.HasNext() {
		seq := fqreader.Next()
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
