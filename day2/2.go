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
	safeCount := 0
	lines := strings.Split(strings.TrimSpace(content), "\n")

	for _, line := range lines {
		parts := strings.Fields(line)
		levels := make([]int, len(parts))

		for i, part := range parts {
			num, _ := strconv.Atoi(part)
			levels[i] = num
		}

		isSafe := false
		isSafe = levelIsSafe(levels)

		if !isSafe {
			for i := 0; i < len(levels); i++ {
				modifiedLevels := []int{}
				modifiedLevels = append(modifiedLevels, levels[:i]...)
				modifiedLevels = append(modifiedLevels, levels[i+1:]...)
				isSafe = levelIsSafe(modifiedLevels)

				if isSafe {
					break
				}
			}
		}

		if isSafe {
			safeCount += 1
		}
	}

	fmt.Println("safeCount:", safeCount)
}

func levelIsSafe(levels []int) bool {
	isIncreasing := true
	isSafe := true
	badLevel := 0

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]

		if diff == 0 {
			badLevel++
			isSafe = false
			break
		}

		if diff < -3 || diff > 3 {
			isSafe = false
			break
		}

		if i == 1 && diff < 0 {
			isIncreasing = false
		} else if (isIncreasing && diff < 0) || (!isIncreasing && diff > 0) {
			isSafe = false
			break
		}
	}

	return isSafe
}
