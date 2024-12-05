package day4

import (
	"aoc/2024/utils"
	"fmt"
	"slices"
	"strings"
)

type Day struct{}

type pos struct {
	x int
	y int
}

const FILE_PATH = "2024/day4/input.txt"

var part1_words = []string{"XMAS", "SAMX"}

func (d Day) Run() {
	fileData := utils.GetFileContent(FILE_PATH)

	const PART1_LETTER = "X"
	const PART2_LETTER = "A"

	matrix := make([][]string, len(fileData))
	part1Pos := []pos{}
	part2Pos := []pos{}

	for y, line := range fileData {
		for x, char := range line {
			c := string(char)
			matrix[y] = append(matrix[y], c)

			if strings.ToUpper(c) == PART1_LETTER {
				part1Pos = append(part1Pos, pos{x: x, y: y})
			} else if strings.ToUpper(c) == PART2_LETTER {
				part2Pos = append(part2Pos, pos{x: x, y: y})
			}
		}
	}

	runPart1(part1Pos, matrix)
	runPart2(part2Pos, matrix)

}

func runPart1(pos []pos, matrix [][]string) {
	count := 0

	for _, coord := range pos {
		countHorizontal(coord, matrix, &count)
		countVertical(coord, matrix, &count)
		countDiagonal(coord, matrix, &count)
	}

	fmt.Printf("Part 1 - 'XMAS' repeats %d times\n", count)
}

func countHorizontal(coord pos, matrix [][]string, count *int) {
	//IF NOT CLOSE TO THE START EDGE
	if coord.x > 2 {
		s := strings.Join(matrix[coord.y][coord.x-3:coord.x+1], "")
		if slices.Contains(part1_words, s) {
			*count++
		}
	}

	//IF NOT CLOSE TO THE END EDGE
	if coord.x < (len(matrix[coord.x]) - 3) {
		s := strings.Join(matrix[coord.y][coord.x:coord.x+4], "")
		if slices.Contains(part1_words, s) {
			*count++
		}
	}
}

func countVertical(coord pos, matrix [][]string, count *int) {
	//IF NOT CLOSE TO THE START EDGE
	if coord.y > 2 {
		s := matrix[coord.y][coord.x] +
			matrix[coord.y-1][coord.x] +
			matrix[coord.y-2][coord.x] +
			matrix[coord.y-3][coord.x]
		if slices.Contains(part1_words, s) {
			*count++
		}
	}

	//IF NOT CLOSE TO THE END EDGE
	if coord.y < (len(matrix) - 3) {
		s := matrix[coord.y][coord.x] +
			matrix[coord.y+1][coord.x] +
			matrix[coord.y+2][coord.x] +
			matrix[coord.y+3][coord.x]
		if slices.Contains(part1_words, s) {
			*count++
		}
	}

}

func countDiagonal(coord pos, matrix [][]string, count *int) {
	//IF NOT CLOSE TO THE BOTTOM EDGE
	if coord.y < (len(matrix) - 3) {
		//IF NOT CLOSE TO THE BOTTOM RIGHT EDGE
		if coord.x < (len(matrix[coord.x]) - 3) {
			s := matrix[coord.y][coord.x] + matrix[coord.y+1][coord.x+1] + matrix[coord.y+2][coord.x+2] + matrix[coord.y+3][coord.x+3]
			if slices.Contains(part1_words, s) {
				*count++
			}
		}

		//IF NOT CLOSE TO THE BOTTOM LEFT EDGE
		if coord.x > 2 {
			s := matrix[coord.y][coord.x] + matrix[coord.y+1][coord.x-1] + matrix[coord.y+2][coord.x-2] + matrix[coord.y+3][coord.x-3]
			if slices.Contains(part1_words, s) {
				*count++
			}
		}
	}

	//IF NOT CLOSE TO THE TOP EDGE
	if coord.y > 2 {
		//IF NOT CLOSE TO THE TOP RIGHT EDGE
		if coord.x > 2 {
			s := matrix[coord.y][coord.x] + matrix[coord.y-1][coord.x-1] + matrix[coord.y-2][coord.x-2] + matrix[coord.y-3][coord.x-3]
			if slices.Contains(part1_words, s) {
				*count++
			}
		}

		//IF NOT CLOSE TO THE TOP LEFT EDGE
		if coord.x < (len(matrix) - 3) {
			s := matrix[coord.y][coord.x] + matrix[coord.y-1][coord.x+1] + matrix[coord.y-2][coord.x+2] + matrix[coord.y-3][coord.x+3]
			if slices.Contains(part1_words, s) {
				*count++
			}
		}
	}
}

func runPart2(pos []pos, matrix [][]string) {
	wordOpt := "MMSS;SMMS;SSMM;MSSM" //possibilities in the edges -- ORDER: TOPL->TOPR->BOTR->BOTL
	count := 0

	for _, coord := range pos {
		if coord.x > 0 && coord.x < len(matrix)-1 && coord.y > 0 && coord.y < len(matrix)-1 {
			val := matrix[coord.y+1][coord.x-1] + matrix[coord.y+1][coord.x+1] + matrix[coord.y-1][coord.x+1] + matrix[coord.y-1][coord.x-1] //This will create the edge of the X
			if strings.Contains(wordOpt, val) {
				count += 1
			}
		}
	}

	fmt.Printf("Part 2 - 'X-MAS' repeats %d times\n", count)
}
