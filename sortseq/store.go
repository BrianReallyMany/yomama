package sortseq

import (
)

type Store struct {
    fileName string
}

func NewStore(fileName string) *Store {
    s := &Store{}

    s.fileName = fileName

    return s
}

func (s *Store) AddSeq(seq SortedSeq) {
}
