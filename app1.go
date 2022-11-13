package main

import (
	"fmt"
	"github.com/mstepan/app1/utils/string_utils"
)

func main() {

	text := "ABABAACABAABABDABABACCABAABABA"
	pattern := "ABABA"

	for i := string_utils.IndexOfKMP(text, pattern, 0); i != -1; i = string_utils.IndexOf(text, pattern, i+1) {
		fmt.Printf("Found at index: %d\n", i)
	}

	fmt.Println("main completed...")
}
