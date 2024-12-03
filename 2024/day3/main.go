package main

import (
	"aoc/2024/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const FILE_PATH = "2024/day3/input.txt"

var REGEXP1 = regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
var REGEXP2 = regexp.MustCompile(`(?s)don't\(\).*?(?:do\(\)|$)`)

func main() {
	mainStart := time.Now()
	lines := utils.GetFileContent(FILE_PATH)

	defer func() {
		fmt.Printf("\nTotal time run -- %.7fs\n", time.Since(mainStart).Seconds())
	}()

	fmt.Println("total sum of mult is ", countMult(lines))
	fmt.Println("total sum of do's mult is ", countDoMult(lines))
}

func countMult(lines []string) int {
	count := 0

	for _, line := range lines {
		finds := REGEXP1.FindAllString(line, -1)
		for _, find := range finds {
			numberString := strings.Replace(strings.Replace(find, "mul(", "", 1), ")", "", 1)
			numbers := strings.Split(numberString, ",")
			num1, _ := strconv.Atoi(numbers[0])
			num2, _ := strconv.Atoi(numbers[1])

			count += (num1 * num2)
		}
	}

	return count
}

func countDoMult(lines []string) int {
	count := 0

	var fullString strings.Builder

	for _, line := range lines {
		fullString.WriteString(line)
	}

	RemoveDont := REGEXP2.ReplaceAllString(fullString.String(), "")
	findsP1 := REGEXP1.FindAllString(RemoveDont, -1)
	for _, v := range findsP1 {
		numberString := strings.Replace(strings.Replace(v, "mul(", "", 1), ")", "", 1)
		numbers := strings.Split(numberString, ",")
		num1, _ := strconv.Atoi(numbers[0])
		num2, _ := strconv.Atoi(numbers[1])

		count += (num1 * num2)
	}

	return count
}
