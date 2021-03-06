package ui

import (
	"bufio"
	"fmt"
	"time"
	"strconv"
	"log"
	"errors"
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
	defer close(ch)

	if len(args) < 1 {
		return
	}

	// TIME
	t0 := time.Now()
	myPath := args[0]

	ch <- "\nVerifying files...\n"
	var fastaFileName, qualFileName string
	fastqFileName, err := getFileByExtension(myPath, "fastq")
	if err != nil {
		ch <- err.Error()
		ch <- "Fastq search failed, checking for fasta and qual files...\n"

		fastaFileName, err = getFileByExtension(myPath, "fasta")
		if err != nil {
			ch <- err.Error()
			ch <- "Fasta search failed, exiting now.\n"
			return
		}

		qualFileName, err = getFileByExtension(myPath, "qual")
		if err != nil {
			ch <- err.Error()
			ch <- "Qual search failed, exiting now.\n"
			return
		}
		ch <- "Fasta and Qual files found!\n"
	} 

	oligoFileName, err := getFileByExtension(myPath, "oligo")
	if err != nil {
		ch <- err.Error()
		ch <- "Oligo search failed, exiting now.\n"
		return
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

	oligofile, err := os.Open(oligoFileName)
	// TODO this is a hack to hold things together until options module is complete
	// then PrepFiles should receive a *SeqSorterOptions as an argument.
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

	t1 := time.Now()
	ch <- fmt.Sprintf("That took %v to run.", t1.Sub(t0))

	ch <- "Reading in sequences..."
	seqcount := 0

	loopStartTime := time.Now()
	lastThousandTime := time.Now()

	for seqReader.HasNext() {
		seq := seqReader.Next()
		//beforeSort := time.Now()
		sortedseq, err := sorter.SortSeq(seq)
		//afterSort := time.Now()
		//ch <- fmt.Sprintf("SortSeq took %v to run", afterSort.Sub(beforeSort))
		if err != nil {
			// Add entry to errorsMap
			// (required type assertion)
			if e, ok := err.(*sortseq.SeqSorterError); ok {
				errorsMap[e.Where] += 1
			}
			seqcount++
			if seqcount % 1000 == 0 {
				ch <- fmt.Sprintf("%d sequences processed ...", seqcount)
				ch <- fmt.Sprintf("Last 1000 sequences took %v to process", time.Now().Sub(lastThousandTime))
				ch <- fmt.Sprintf("So far we've processed %d sequences in %v\n", seqcount, time.Now().Sub(loopStartTime))
				lastThousandTime = time.Now()
			}
			continue
		}
		err = store.AddSeq(sortedseq)
		if err != nil {
			ch <- "PrepFiles: error storing seq--"
			ch <- sortedseq.ToString()
		}
		seqcount++
		if seqcount % 1000 == 0 {
			ch <- strconv.Itoa(seqcount) + " sequences processed ..."
			ch <- fmt.Sprintf("Last 1000 sequences took %v to process", time.Now().Sub(lastThousandTime))
			ch <- fmt.Sprintf("So far we've processed %d sequences in %v\n", seqcount, time.Now().Sub(loopStartTime))
			lastThousandTime = time.Now()
		}
	}
	ch <- "...done.\n"
	ch <- "Error summary:"
	ch <- "Type of error\tCount"
	for k, v := range errorsMap {
		ch <- k + "\t" + strconv.Itoa(v) + "\n"
	}
}

func (c *MamaController) System(args []string) string {
	out, err := exec.Command(args[0], args[1:]...).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func getFileByExtension(path, extension string) (string, error) {
	files, err := filepath.Glob(path + "/*." + extension)
	if err != nil {
		return "", err
	}
	if len(files) == 0 {
		return "", errors.New("No files found with extension ." + extension)
	} else if len(files) > 1 {
		return "", errors.New("Multiple files found with extension ." + extension)
	}
	return files[0], nil
}

