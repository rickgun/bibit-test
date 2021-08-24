package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	str := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	fmt.Println(anagramGrouping(str))
}

func anagramGrouping(sliceStr []string) (result [][]string) {
	mapString := make(map[string][]string)
	for _, s := range sliceStr {
		key := sortString(s)
		mapString[key] = append(mapString[key], s)
	}

	for i := range mapString {
		result = append(result, mapString[i])
	}

	return result
}

func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)

	return strings.Join(s, "")
}
