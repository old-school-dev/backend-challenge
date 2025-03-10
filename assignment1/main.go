package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func DFS(graph, memo [][]int, row, col int) int {
	if memo[row][col] != -1 {
		return memo[row][col]
	}
	// catch last row (break point)
	if row == len(graph)-1 {
		memo[row][col] = graph[row][col]
		return memo[row][col]
	}
	nextRow := row + 1

	leftNode := col
	rightNode := col + 1

	leftRoute := DFS(graph, memo, nextRow, leftNode)
	rightRoute := DFS(graph, memo, nextRow, rightNode)

	memo[row][col] = graph[row][col] + max(leftRoute, rightRoute)
	return memo[row][col]
}

func Solve(graph [][]int) {
	memo := make([][]int, len(graph))
	for i := range graph {
		memo[i] = make([]int, len(graph[i]))
		for j := range graph[i] {
			memo[i][j] = -1
		}
	}
	result := DFS(graph, memo, 0, 0)
	fmt.Println(result)
}

func readFile(filename string) (graph [][]int) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	err = json.Unmarshal(jsonData, &graph)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil
	}
	return
}

func main() {
	// easy
	// graph := readFile("easy.json")
	graph := readFile("hard.json")
	Solve(graph)
}
