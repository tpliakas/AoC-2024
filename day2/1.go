package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	content, err := utils.ReadFileContents("day2/input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	lines := strings.Split(strings.TrimSpace(content), "\n")
	safeCount := 0

	for _, line := range lines {
		parts := strings.Fields(line)
		report := make([]int, len(parts))
		for i, part := range parts {
			report[i], err = strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return
			}
		}

		if utils.IsSafeReport(report) {
			safeCount++
		}
	}

	fmt.Printf("Number of safe reports: %d\n", safeCount)
}
