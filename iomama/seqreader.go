package iomama

import (
	. "github.com/BrianReallyMany/yomama/seq"
)

type SeqReader interface {
	HasNext() bool
	Next() Seq
}
