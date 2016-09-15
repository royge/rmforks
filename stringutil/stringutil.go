package stringutil

import (
	"sort"
)

func Contains(hayStack []string, target string) bool {
	sort.Strings(hayStack)
	i := sort.SearchStrings(hayStack, target)
	return i < len(hayStack) && hayStack[i] == target
}
