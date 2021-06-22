package maps

import (
	"sort"
	"strings"
)

func SortedValues(m map[int]string) string {
	values := make([]string, 0, len(m))
	keys := make([]int, 0, len(m))
	for i, _ := range m {
		keys = append(keys, i)
	}
	sort.Ints(keys)
	for _, k := range keys {
		values = append(values, m[k])
	}
	concatenate := strings.Join(values[:], "")
	return concatenate
}
