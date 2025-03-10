package main

import "fmt"

func Solve(encodedString string) string {
	n := len(encodedString) + 1
	result := make([]int, n)
	resultTemp := make([]int, n)
	minSum := 9*len(encodedString) + 1
	var backtrack func(index int)

	backtrack = func(index int) {
		if index == n {
			sum := 0
			for _, v := range resultTemp {
				sum += v
			}
			if sum < minSum {
				minSum = sum
				copy(result, resultTemp)
			}
			return
		}

		for i := range 10 {
			if index == 0 || (encodedString[index-1] == 'L' && resultTemp[index-1] > i) ||
				(encodedString[index-1] == 'R' && resultTemp[index-1] < i) ||
				(encodedString[index-1] == '=' && resultTemp[index-1] == i) {
				resultTemp[index] = i
				backtrack(index + 1)
			}
		}
	}
	backtrack(0)
	answer := ""
	for _, r := range result {
		answer += fmt.Sprintf("%d", r)
	}
	return answer
}

func main() {
	encodedString := "LLRR=" 
	Solve(encodedString)
}
