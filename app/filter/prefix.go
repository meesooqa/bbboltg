package filter

import (
	"strings"
)

type PrefixFilter struct {
	prefix string
}

func newPrefixFilter(prefix string) *PrefixFilter {
	return &PrefixFilter{prefix: prefix}
}

func (f *PrefixFilter) IsValid(s string) bool {
	return strings.HasPrefix(s, f.prefix)
}
