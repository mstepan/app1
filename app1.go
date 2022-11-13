package main

import (
	"fmt"
	"github.com/mstepan/app1/utils/number_utils"
	"github.com/mstepan/app1/utils/string_utils"
)

func main() {

	for i := -5; i < 10; i++ {

		res, err := number_utils.Fibonacci(i)

		if err == nil {
			fmt.Printf("fibonacci(%d) = %d\n", i, res)
		} else {
			fmt.Printf("error: %v\n", err)
		}

	}

	str := "aB"
	pattern := "A"

	index := string_utils.IndexOf(str, pattern)

	fmt.Printf("str: %s, pattern: %s, index: %d \n", str, pattern, index)

	fmt.Println("main completed...")
}
