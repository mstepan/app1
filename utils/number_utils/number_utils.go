package number_utils

import "fmt"

func Fac(val int) (int, error) {

	if val < 0 {
		return -1, fmt.Errorf("can't calculate factorial from negative value: %d", val)
	}

	res := 1

	for curValue := 2; curValue <= val; curValue++ {
		res *= curValue
	}

	return res, nil
}

func Fibonacci(index int) (int, error) {

	if index < 0 {
		return -1, fmt.Errorf("can't calculate fibonacci for negative index: %d", index)
	}

	first := 0
	second := 1

	for i := 0; i <= index; i++ {
		first, second = second, first+second
	}

	return first, nil

}
