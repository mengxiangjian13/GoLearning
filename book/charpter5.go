package main

import (
	"bytes"
)

func main() {

}

//
func isPalindrome(input string) bool {
	ret := true
	length := len(input)
	count := length / 2
	for i := 0; i < count; i++ {
		if input[i] != input[length-i] {
			ret = false
			break
		}
	}
	return ret
}

func commonPrefix(texts []string) string {
	slices := make([][]rune, len(texts))
	for i, v := range texts {
		slices[i] = []rune(v)
	}
	if len(slices) == 0 {
		return ""
	}
	var common = bytes.Buffer
FINISH:
	for i := 0; i < len(slices[0]); i++ {
		char := slices[0][i]
		for j := 1; j < len(slices); j++ {
			if i >= len(slices[j]) || slices[j][i] != char {
				break FINISH
			}
		}
		common.WriteRune(char)
	}
	return common.String()
}
