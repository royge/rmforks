package stringutil

import (
	"sort"
)

// Contains check if a string is present in a slice.
func Contains(l []string, s string) bool {
	sort.Strings(l)
	i := sort.SearchStrings(l, s)
	return i < len(l) && l[i] == s
}
