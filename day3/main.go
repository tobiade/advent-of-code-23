package main

import (
	"fmt"
	"strings"

	"github.com/tobiade/advent-of-code-23/utils"
)

type Coord struct {
	row int
	col int
}

func main() {
	part1()
	part2()
}

func part1() {
	lines := utils.ReadLines("input.txt")
	numRows := len(lines)
	numCols := len(lines[0])
	grid := make([][]string, numRows)
	seen := make([][]bool, numRows)
	for i := range grid {
		grid[i] = make([]string, numCols)
		seen[i] = make([]bool, numCols)
	}

	for idx, line := range lines {
		chars := strings.Split(line, "")
		if len(chars) != numCols {
			panic(fmt.Sprintf("invalid number of colums - expected %d got %d", numCols, len(chars)))
		}
		grid[idx] = chars
	}

	symbols := make([]Coord, 0)
	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			char := grid[r][c]
			if char != "." && !isDigit(char) {
				symbols = append(symbols, Coord{row: r, col: c})
			}
		}
	}

	directions := [][]int{
		{0, -1},  // left
		{0, 1},   // right
		{-1, 0},  // up
		{1, 0},   // down
		{-1, -1}, // upper left
		{-1, 1},  // upper right
		{1, -1},  // lower left
		{1, 1},   // lower right
	}
	sum := 0
	for _, symbol := range symbols {
		seen[symbol.row][symbol.col] = true
		for _, d := range directions {
			nR := symbol.row + d[0]
			nC := symbol.col + d[1]
			if nR < 0 || nR >= numRows || nC < 0 || nC >= numCols || seen[nR][nC] {
				continue
			}

			if isDigit(grid[nR][nC]) {
				sum += markAndScan(nR, nC, grid, seen)
			}
		}
	}

	fmt.Println("part1: ", sum)
}

func isDigit(s string) bool {
	return rune(s[0]) >= '0' && rune(s[0]) <= '9'
}

func markAndScan(row int, col int, grid [][]string, seen [][]bool) int {
	leftEnd := col - 1
	rightStart := col + 1

	var leftPart string
	for leftEnd >= 0 {
		if isDigit(grid[row][leftEnd]) {
			leftPart += grid[row][leftEnd]
			seen[row][leftEnd] = true
			leftEnd--
		} else {
			break
		}
	}
	leftPart = utils.Reverse(leftPart)

	var rightPart string
	for rightStart < len(grid[0]) {
		if isDigit(grid[row][rightStart]) {
			rightPart += grid[row][rightStart]
			seen[row][rightStart] = true
			rightStart++
		} else {
			break
		}
	}

	numStr := leftPart + grid[row][col] + rightPart
	return utils.MustAtoi(numStr)
}

func part2() {
	lines := utils.ReadLines("input.txt")
	numRows := len(lines)
	numCols := len(lines[0])
	grid := make([][]string, numRows)
	seen := make([][]bool, numRows)
	for i := range grid {
		grid[i] = make([]string, numCols)
		seen[i] = make([]bool, numCols)
	}

	for idx, line := range lines {
		chars := strings.Split(line, "")
		if len(chars) != numCols {
			panic(fmt.Sprintf("invalid number of colums - expected %d got %d", numCols, len(chars)))
		}
		grid[idx] = chars
	}

	symbols := make([]Coord, 0)
	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			char := grid[r][c]
			if char == "*" {
				symbols = append(symbols, Coord{row: r, col: c})
			}
		}
	}

	directions := [][]int{
		{0, -1},  // left
		{0, 1},   // right
		{-1, 0},  // up
		{1, 0},   // down
		{-1, -1}, // upper left
		{-1, 1},  // upper right
		{1, -1},  // lower left
		{1, 1},   // lower right
	}
	sum := 0
	for _, symbol := range symbols {
		seen[symbol.row][symbol.col] = true
		nums := make([]int, 0)
		for _, d := range directions {
			nR := symbol.row + d[0]
			nC := symbol.col + d[1]
			if nR < 0 || nR >= numRows || nC < 0 || nC >= numCols || seen[nR][nC] {
				continue
			}

			if isDigit(grid[nR][nC]) {
				nums = append(nums, markAndScan(nR, nC, grid, seen))
			}
		}

		if len(nums) == 2 {
			sum += nums[0] * nums[1]
		}
	}

	fmt.Println("part2: ", sum)
}
