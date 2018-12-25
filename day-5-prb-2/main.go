package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	maxInt = int(^uint(0) >> 1)
	minInt = -maxInt - 1
)

func main() {
	f, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var removeResult string
	dd := make(map[string]int)
	mapOfUniqueChars := getUniqueChars(string(f))
	for key := range mapOfUniqueChars {
		removeResult = removeChars(rune(key), string(f))
		dda := recursive(removeResult)
		dd[string(key)] = len(dda)
	}
	fmt.Println(dd)
	fmt.Println("The minumum length is ", min(dd))
}

func min(input map[string]int) int {
	min := maxInt
	for _, d := range input {
		if d < min {
			min = d
		}
	}
	return min
}

func removeChars(a rune, source string) string {
	smallLetter := getSmallLetter(a)
	var result []rune
	for _, char := range source {
		if char == a || char == smallLetter {
			continue
		}
		result = append(result, char)
	}
	return string(result)
}

func getSmallLetter(a rune) rune {
	const diff = 'a' - 'A'
	return a + diff
}

func getUniqueChars(input string) map[rune]int {

	dd := make(map[rune]int)
	for _, char := range input {
		capital := strings.ToUpper(string(char))
		for _, c := range capital {
			runed := rune(c)
			dd[runed]++
		}

	}
	return dd
}

func recursive(s string) string {
	ok := true
	for ok {
		s, ok = step(s)
	}
	return s
}

func step(s string) (string, bool) {
	for i := 0; i < len(s)-1; i++ {
		if opposite(s[i], s[i+1]) {
			return s[:i] + s[i+2:], true
		}
	}
	return s, false
}

func opposite(a, b byte) bool {
	const diff = 'a' - 'A'
	return a+diff == b || b+diff == a
}
