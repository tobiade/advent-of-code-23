package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/tobiade/advent-of-code-23/utils"
)

var regexPttrn = regexp.MustCompile("[1-9]|one|two|three|four|five|six|seven|eight|nine")
var regexPttrnRvrs = regexp.MustCompile("[1-9]|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin") // LOL

var digits = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"eno":   "1",
	"owt":   "2",
	"eerht": "3",
	"ruof":  "4",
	"evif":  "5",
	"xis":   "6",
	"neves": "7",
	"thgie": "8",
	"enin":  "9",
}

func main() {
	part1()
	part2()
}

func part1() {
	f, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	doc := strings.Split(string(f), "\n")

	sum := 0
	for _, val := range doc {
		start := 0
		end := len(val) - 1

		for start < len(val) && !isNum(rune(val[start])) {
			start++
		}

		for end >= 0 && !isNum(rune(val[end])) {
			end--
		}

		if start <= end {
			numStr := string(val[start]) + string(val[end])
			num, err := strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}
			sum += num
		}
	}

	fmt.Println(sum)
}

func isNum(r rune) bool {
	return r >= '0' && r <= '9'
}

func part2() {

	f, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	doc := strings.Split(string(f), "\n")

	sum := 0
	for _, val := range doc {
		sum += getNum(val)
	}
	fmt.Println(sum)
}

func getNum(val string) int {
	s := regexPttrn.FindAllString(val, -1)
	if s == nil {
		panic("no match found")
	}

	reversed := utils.Reverse(val)
	x := regexPttrnRvrs.FindAllString(reversed, -1)
	if x == nil {
		panic("no match found")
	}
	firstDigit := s[0]
	secondDigit := x[0]

	if _, ok := digits[firstDigit]; ok {
		firstDigit = digits[firstDigit]
	}

	if _, ok := digits[secondDigit]; ok {
		secondDigit = digits[secondDigit]
	}

	numStr := firstDigit + secondDigit
	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic(err)
	}
	return num
}
