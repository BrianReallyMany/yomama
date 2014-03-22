package sequtil


var m = map[byte][]byte{
	byte('a'): []byte{'a'},
	byte('c'): []byte{'c'},
	byte('g'): []byte{'g'},
	byte('t'): []byte{'t', 'u'},
	byte('u'): []byte{'t', 'u'},
	byte('r'): []byte{'a', 'g'},
	byte('y'): []byte{'c', 't'},
	byte('s'): []byte{'g', 'c'},
	byte('w'): []byte{'a', 't'},
	byte('k'): []byte{'g', 't'},
	byte('m'): []byte{'a', 'c'},
	byte('b'): []byte{'c', 'g', 't'},
	byte('d'): []byte{'a', 'g', 't'},
	byte('h'): []byte{'a', 'c', 't'},
	byte('v'): []byte{'a', 'c', 'g'},
	byte('n'): []byte{'n', 'a', 'c', 'g', 't'},
}

// Assumes specified order for arguments; violators will be prosecuted
func MatchBase(oligobase, rawbase byte) bool {
	matches, _ := m[oligobase]
	if sliceContainsBase(matches, rawbase) {
		return true
	}
	return false
}


func sliceContainsBase(baseSlice []byte, rawbase byte) bool {
	for _, base := range baseSlice {
		if rawbase == base {
			return true
		}
	}
	return false
}

