package main

import (
	"aoc/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

type pair struct {
	sum        int
	distance   int
	similarity int
}

func main() {
	mainStart := time.Now()
	lines := utils.GetFileContent("2024/day1/input.txt")

	defer func() {
		fmt.Printf("\nTotal time run -- %.7fs\n", time.Since(mainStart).Seconds())
	}()

	leftArray := []int{}
	rightArray := []int{}

	for _, line := range lines {
		split := strings.Split(line, "   ")
		num1, _ := strconv.Atoi(split[0])
		num2, _ := strconv.Atoi(split[1])

		leftArray = append(leftArray, num1)
		rightArray = append(rightArray, num2)
	}

	slices.Sort(leftArray)
	slices.Sort(rightArray)

	fmt.Println("--------- RUNNING DAY 1 ---------")

	start := time.Now()
	result := pairUp(leftArray, rightArray)
	fmt.Printf("Time consumed by pairUp: %.7fs\n", time.Since(start).Seconds())

	distance := findTotalOfValue(result, func(p pair) int { return p.distance })

	fmt.Printf("Total distance between: %d\n", distance)

	similarity := findTotalOfValue(result, func(p pair) int { return p.similarity })
	fmt.Println("Similarity score:", similarity)

}

func pairUp(slice1 []int, slice2 []int) []pair {
	n := len(slice1)

	if n != len(slice2) {
		return []pair{}
	}

	pairArray := []pair{}

	for i, val := range slice1 {
		val2 := slice2[i]
		similarity := findSimilarity(val, slice2)
		pairArray = append(pairArray, pair{sum: val + val2, distance: utils.Abs(val - val2), similarity: similarity})
	}

	return pairArray
}

func findTotalOfValue(pairArray []pair, getFieldValue func(pair) int) int {
	total := 0

	for _, value := range pairArray {
		total += getFieldValue(value)
	}

	return total
}

func findSimilarity(val int, array []int) int {
	count := 0
	for _, value := range array {
		if value == val {
			count += 1
		}
	}

	return val * count
}
