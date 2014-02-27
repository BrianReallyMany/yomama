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

func MatchBase(base1, base2 string) bool {
	matches := m[strings.ToLower(base1)]
	return strings.Contains(matches, strings.ToLower(base2))
}


