package main

import (
	"fmt"
	"sort"
)

func main() {
	h := map[string]int{
		"something": 10,
		"yo":        20,
		"blah":      20,
	}

	type hv struct {
		Key   string
		Value int
	}

	var hv1 []hv
	for k, v := range h {
		hv1 = append(hv1, hv{k, v})
	}

	sort.Slice(hv1, func(i, j int) bool {

		return hv1[i].Value < hv1[j].Value
	})

	for _, hv := range hv1 {
		fmt.Printf("%s, %d\n", hv.Key, hv.Value)
	}

}
