package string_utils

/*
IndexOf - find substring inside string using naive string matching algorithm.
time: O(N+M), where N - len(str), M - len(pattern)
space: O(1)
*/
func IndexOf(str string, pattern string) int {

	if len(pattern) > len(str) {
		return -1
	}

	for i := 0; i <= len(str)-len(pattern); i++ {
		j := 0

		for j < len(pattern) && str[i+j] == pattern[j] {
			j++
		}

		if j == len(pattern) {
			return i
		}
	}

	return -1
}
