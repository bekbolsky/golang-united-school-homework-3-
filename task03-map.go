package homework

import (
	"sort"
)

// sortMapValues returns map values sorted in order of increasing keys
func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, 0, len(input))
	for k := range input {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		result = append(result, input[k])
	}
	return
}
