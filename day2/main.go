package main

import (
	"fmt"
	"strings"

	"github.com/tobiade/advent-of-code-23/utils"
)

const (
	totalRed   = 12
	totalGreen = 13
	totalBlue  = 14
)

type Game struct {
	gameNum int
	sets    [][]Set
}

type Set struct {
	colour   string
	numCubes int
}

// Not my best work but it works
func main() {
	part1()
	part2()
}

func part1() {
	games := parse(utils.ReadLines("input.txt"))

	sum := 0
	for _, game := range games {
		add := true
		for _, set := range game.sets {
			m := make(map[string]int)
			for _, s := range set {
				m[s.colour] = m[s.colour] + s.numCubes
			}
			if m["red"] > totalRed || m["green"] > totalGreen || m["blue"] > totalBlue {
				add = false
			}
		}
		if add {
			sum += game.gameNum
		}
	}

	fmt.Println("part1: ", sum)
}

func parse(games []string) []Game {
	res := make([]Game, 0)
	for _, game := range games {
		gameAndSets := strings.Split(game, ":")
		gameNum := strings.Split(gameAndSets[0], " ")[1]
		sets := strings.Split(gameAndSets[1], ";")
		parsedSets := parseSet(sets)
		gameObj := Game{gameNum: utils.MustAtoi(gameNum), sets: parsedSets}
		res = append(res, gameObj)
	}
	return res
}

func parseSet(sets []string) [][]Set {
	res := make([][]Set, 0)
	for _, set := range sets {
		s := make([]Set, 0)
		picks := strings.Split(set, ",")
		for _, pick := range picks {
			pick = strings.TrimSpace(pick)
			numberAndColour := strings.Split(pick, " ")
			s = append(s, Set{numCubes: utils.MustAtoi(numberAndColour[0]), colour: numberAndColour[1]})
		}
		res = append(res, s)
	}
	return res
}

func part2() {
	games := parse(utils.ReadLines("input.txt"))

	sum := 0
	for _, game := range games {
		maxRed := 0
		maxGreen := 0
		maxBlue := 0
		for _, set := range game.sets {
			m := make(map[string]int)
			for _, s := range set {
				m[s.colour] = m[s.colour] + s.numCubes
			}
			maxRed = max(maxRed, m["red"])
			maxGreen = max(maxGreen, m["green"])
			maxBlue = max(maxBlue, m["blue"])
		}
		power := maxRed * maxGreen * maxBlue
		sum += power
	}

	fmt.Println("part2: ", sum)
}
