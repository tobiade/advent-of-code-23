package utils

import (
	"os"
	"strconv"
	"strings"
)

func ReadLines(path string) []string {
	f, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(f), "\n")
}

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
