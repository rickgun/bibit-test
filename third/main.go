package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "There is new project (design in progress) coming up!"
	fmt.Println(findFirstStringInBracket(str))
}

func findFirstStringInBracket(str string) string {
	if len(str) > 0 {
		first := strings.Index(str, "(")
		last := strings.Index(str, ")")

		if first >= 0 && last >= 0 {
			return str[first+1 : last]
		}
	}
	return ""
}
