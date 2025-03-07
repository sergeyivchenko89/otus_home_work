package hw03frequencyanalysis

import (
	"regexp"
	"sort"
)

var validID = regexp.MustCompile(`[^ \n\t]+`)

func Top10(str string) []string {
	if len(str) == 0 {
		return []string{}
	}

	source := validID.FindAllString(str, -1)
	dest := make([]string, 0)
	dict := make(map[string]uint)
	for _, v := range source {
		if dict[v] == 0 {
			dest = append(dest, v)
		}
		dict[v]++
	}
	sort.Slice(dest, func(i, j int) bool {
		return dict[dest[i]] > dict[dest[j]] || dict[dest[i]] == dict[dest[j]] && dest[i] < dest[j]
	})

	return dest[0:10]
}
