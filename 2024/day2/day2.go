package day2

import (
	"aoc/2024/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Day2 struct{}

func (d Day2) Run() {
	lines := utils.GetFileContent("2024/day2/input.txt")

	count, countDamp := 0, 0

	for _, line := range lines {
		levels := strings.Split(line, " ")

		isSafe := checkReportSafe(levels)
		if isSafe {
			count++
		} else if checkDampener(levels) {
			countDamp++
		}
	}

	countDamp += count

	fmt.Printf("Exists %d reports safe\n", count)
	fmt.Printf("Exists %d reports safe with dampener\n", countDamp)
}

func checkReportSafe(levels []string) bool {
	firstLevel, _ := strconv.Atoi(levels[0])
	currentLevel, _ := strconv.Atoi(levels[1])
	isDecreasing := firstLevel > currentLevel

	if !isValidDiff(firstLevel, currentLevel) {
		return false
	}

	for _, level := range levels[2:] {
		levelConv, _ := strconv.Atoi(level)

		if !isValidDiff(levelConv, currentLevel) || (isDecreasing && levelConv > currentLevel) || (!isDecreasing && levelConv < currentLevel) {
			return false
		}

		currentLevel = levelConv
	}

	return true
}

func isValidDiff(num1 int, num2 int) bool {
	diff := utils.Abs(num1 - num2)
	if diff > 3 || diff == 0 {
		return false
	}

	return true
}

func checkDampener(levels []string) bool {
	for i := 0; i < len(levels); i++ {
		copyLevels := make([]string, len(levels))
		copy(copyLevels, levels)
		copyLevels = slices.Delete(copyLevels, i, i+1)
		if checkReportSafe(copyLevels) {
			return true
		}
	}

	return false
}
