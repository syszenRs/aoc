package main

import (
	"aoc/2024/day1"
	"aoc/2024/day2"
	"aoc/2024/day3"
	"aoc/2024/utils"
	"fmt"
	"os"
	"strconv"
	"time"
)

var dayFunc = map[int]utils.DayInterface{
	1: day1.Day1{},
	2: day2.Day2{},
	3: day3.Day3{},
}

var DAYS_COMPLETED = len(dayFunc)

func main() {
	fmt.Println(`
	-------------------------------
	Welcome to Advent of Code 2024!
	-------------------------------
	`)

	var daySelected int

	cliArgs := os.Args
	if len(cliArgs) > 1 {
		daySelected, _ = strconv.Atoi(os.Args[1])
	}

	if !isValidDaySelected(daySelected) {
		for {
			fmt.Printf("Select the day to run (we have %d days available): ", DAYS_COMPLETED)
			fmt.Scanln(&daySelected)
			if isValidDaySelected(daySelected) {
				break
			} else {
				fmt.Println("invalid day was selected, try again..")
			}
		}
	}

	mainStart := time.Now()

	defer func() {
		fmt.Printf("\nPuzzle run finished!\nTotal time run -- %.7fs\n", time.Since(mainStart).Seconds())
	}()

	fmt.Printf("\nRunning day %d puzzle\n", daySelected)

	dayFunc[daySelected].Run()
}

func isValidDaySelected(day int) bool {
	return day >= DAYS_COMPLETED && day <= DAYS_COMPLETED
}
