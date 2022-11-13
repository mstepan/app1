package string_utils

/*
IndexOf - find substring inside string using naive string matching algorithm.
time: O(N+M), where N - len(text), M - len(pattern)
space: O(1)
*/
func IndexOf(text string, pattern string, from int) int {
	if len(pattern) > len(text) {
		return -1
	}

	for i := from; i <= len(text)-len(pattern); i++ {
		j := 0

		for j < len(pattern) && text[i+j] == pattern[j] {
			j++
		}

		if j == len(pattern) {
			return i
		}
	}

	return -1
}

func IndexOfFromBeginning(text string, pattern string) int {
	return IndexOf(text, pattern, 0)
}

/*
IndexOfKMP find index of 'pattern' string matching with 'text' string.
Use Knuth-Morris-Pratt algorithm
time: O(N+M), where N - len(text), M - len(pattern)
space: O(M)
*/
func IndexOfKMP(text string, pattern string, from int) int {

	textLength := len(text)
	patternLength := len(pattern)

	if patternLength > textLength-from {
		return -1
	}

	prefixTable := prefixTable(pattern)

	j := 0
	for i := from; i < textLength; {
		if j == patternLength {
			return i - len(pattern)
		}

		if text[i] == pattern[j] {
			i++
			j++
		} else {
			if j != 0 {
				j = prefixTable[j-1]
			}
			if text[i] == pattern[j] {
				i++
				j++
			} else {
				i++
				j = 0
			}
		}
	}

	// 'pattern' string found at the end of 'text' string
	if j == patternLength {
		return textLength - patternLength
	}

	return -1
}

/*
*
Create prefix table for KMP string matching algorithm.
*/
func prefixTable(str string) []int {
	prefix := make([]int, len(str))

	for i := 1; i < len(str); i++ {

		ch := str[i]
		index := i - 1

		for index >= 0 {
			lastCh := str[prefix[index]]

			if ch == lastCh {
				prefix[i] = prefix[index] + 1
				break
			} else {
				index = prefix[index] - 1
			}
		}

		if index < 0 {
			prefix[i] = 0
		}
	}

	return prefix
}
