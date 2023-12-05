package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/tobiade/advent-of-code-23/utils"
)

var re = regexp.MustCompile(`\s+`)

type Card struct {
	cardNum      int
	winningNums  []int
	otherNums    []int
	matchingNums int
}

func main() {
	cards := part1()
	part2(cards)
}

func part1() []Card {
	lines := utils.ReadLines("input.txt")
	cards := parse(lines)
	res := make([]Card, 0)

	sum := int64(0)
	for _, card := range cards {
		m := make(map[int]int)
		for _, n := range card.winningNums {
			m[n] += 1
		}

		count := 0
		for _, n := range card.otherNums {
			if _, ok := m[n]; ok {
				count++
				m[n] -= 1
				if m[n] == 0 {
					delete(m, n)
				}
			}
		}

		points := math.Pow(2, float64(count-1))
		sum += int64(points)

		// For part 2
		card.matchingNums = count
		res = append(res, card)
	}
	fmt.Println("part1: ", sum)

	return res
}

func parse(lines []string) []Card {
	cards := make([]Card, 0)
	for _, line := range lines {
		line = re.ReplaceAllString(line, " ")
		cardAndNums := strings.Split(line, ":")
		cardNumStr := strings.Split(cardAndNums[0], " ")[1]
		cardNum := utils.MustAtoi(cardNumStr)

		nums := strings.Split(cardAndNums[1], "|")
		winningNumStrs := strings.Split(strings.TrimSpace(nums[0]), " ")

		winningNums := make([]int, 0)
		for _, n := range winningNumStrs {
			winningNums = append(winningNums, utils.MustAtoi(n))
		}

		otherNumStrs := strings.Split(strings.TrimSpace(nums[1]), " ")
		otherNums := make([]int, 0)
		for _, n := range otherNumStrs {
			otherNums = append(otherNums, utils.MustAtoi(n))
		}

		cards = append(cards, Card{cardNum: cardNum, winningNums: winningNums, otherNums: otherNums})
	}
	return cards
}

func part2(cards []Card) {
	freq := make(map[int]int)
	for _, card := range cards {
		freq[card.cardNum] += 1
	}

	count := 0
	for _, card := range cards {
		for i := 0; i < freq[card.cardNum]; i++ {
			count++
			for i := 1; i <= card.matchingNums; i++ {
				nextCardNum := card.cardNum + i
				freq[nextCardNum] += 1

			}
		}
	}

	fmt.Println("part2: ", count)
}
