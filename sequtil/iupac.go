package sequtil

import (
	"strings"
)

var m = map[string]string{
	"a": "arwmdhvn",
	"c": "cysmbhvn",
	"g": "grskbdvn",
	"t": "tuywkbdhn",
	"u": "tuywkbdhn",
	"r": "rag",
	"y": "yct",
	"s": "sgc",
	"w": "wat",
	"k": "kgt",
	"m": "mac",
	"b": "bcgt",
	"d": "dagt",
	"h": "hact",
	"v": "vacg",
	"n": "nacgturyswkmbdhv",
}

func MatchBase(base1, base2 string) bool {
	matches := m[strings.ToLower(base1)]
	return strings.Contains(matches, strings.ToLower(base2))
}


