package day3

import (
	"aoc/2024/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Day3 struct{}

const FILE_PATH = "2024/day3/input.txt"

var REGEXP1 = regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
var REGEXP2 = regexp.MustCompile(`(?s)don't\(\).*?(?:do\(\)|$)`)

func (d Day3) Run() {
	RunPart1()
	RunPart2()
}

func RunPart1() {
	lines := utils.GetFileContent(FILE_PATH)
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

	fmt.Println("total sum of mult is ", count)
}

func RunPart2() {
	lines := utils.GetFileContent(FILE_PATH)
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

	fmt.Println("total sum of do's mult is ", count)
}
