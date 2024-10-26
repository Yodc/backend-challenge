package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxPath(triangle [][]int) int {
	dp := make([][]int, len(triangle))
	for i := range triangle {
		dp[i] = make([]int, len(triangle[i]))
		copy(dp[i], triangle[i])
	}

	for i := len(dp) - 2; i >= 0; i-- {
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] += max(dp[i+1][j], dp[i+1][j+1])
		}
	}

	return dp[0][0]
}

func main() {

	jsonFile, err := os.Open("test.json")
	if err != nil {
		log.Fatalf("Failed to open JSON file: %s", err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %s", err)
	}

	var data [][]int
	if err := json.Unmarshal(byteValue, &data); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %s", err)
	}

	json.Unmarshal([]byte(byteValue), &data)
	maxPathSum := findMaxPath(data)
	fmt.Printf("เส้นทางที่มีค่ามากที่สุดคือ: %d\n", maxPathSum)
}
