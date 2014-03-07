package sequtil

import (
	"strings"
)

var m = map[string]string{
	"a": "a",
	"c": "c",
	"g": "g",
	"t": "tu",
	"u": "tu",
	"r": "ag",
	"y": "ct",
	"s": "gc",
	"w": "at",
	"k": "gt",
	"m": "ac",
	"b": "cgt",
	"d": "agt",
	"h": "act",
	"v": "acg",
	"n": "nacgt",
}

// Assumes specified order for arguments; violators will be prosecuted
func MatchBase(oligobase, rawbase string) bool {
	matches := m[strings.ToLower(oligobase)]
	return strings.Contains(matches, strings.ToLower(rawbase))
}


