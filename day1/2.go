package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	contents, err := utils.ReadFileContents("day1/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	var leftList, rightList []int
	lines := strings.Split(contents, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		numbers := strings.Fields(line)
		if len(numbers) != 2 {
			fmt.Println("Invalid line:", line)
			continue
		}

		num1, err1 := strconv.Atoi(numbers[0])
		num2, err2 := strconv.Atoi(numbers[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Error parsing numbers on line:", line)
			continue
		}

		leftList = append(leftList, num1)
		rightList = append(rightList, num2)
	}

	rightCount := make(map[int]int)
	for _, num := range rightList {
		rightCount[num]++
	}

	var similarityScore int
	for _, num := range leftList {
		count, exists := rightCount[num]
		if exists {
			similarityScore += num * count
		}
	}

	fmt.Println("Total similarity score:", similarityScore)
}
