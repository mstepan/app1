package main

import "fmt"

func main() {

	printNFibonacci(20)

	for i := 0; i <= 20; i += 1 {
		fmt.Printf("factorial(%d) = %d\n", i, factorial(i))
	}

	fmt.Println("main completed...")
}

func printNFibonacci(n int) {
	var first = 1
	var second = 1

	for i := 0; i < n; i += 1 {
		fmt.Printf("fib(%d) = %d\n", i, first)
		first, second = second, first+second
	}

}

func factorial(value int) int {
	if value == 0 {
		return 1
	}
	if value < 2 {
		return value
	}
	return value * factorial(value-1)
}
