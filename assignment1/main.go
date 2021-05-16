package main

import (
	"fmt"
)

func balancingBracket(str string) bool {
	arr := make([]int, 6)
	chars := []rune(str)

	if str == "" {
		return false
	}

	for i := 0; i < len(chars); i++ {
		if chars[i] == '(' {
			arr[0]++
		} else if chars[i] == ')' {
			arr[0]++
		} else if chars[i] == '[' {
			arr[1]++
		} else if chars[i] == ']' {
			arr[1]++
		} else if chars[i] == '{' {
			arr[2]++
		} else if chars[i] == '}' {
			arr[2]++
		}
	}

	for j := 0; j < len(arr); j++ {
		if arr[j]%2 != 0 {
			return false
		}

	}
	return true
}

func main() {

	str := "[(]))"
	fmt.Println(balancingBracket(str))

}
