package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	sumOfTwo := 0
	sumOfThree := 0
	for s.Scan() {
		currentText := s.Text()
		charCountMap := charCount(currentText)
		countOfTwo, countOfThree := mapper(charCountMap)
		sumOfTwo += countOfTwo
		sumOfThree += countOfThree
	}

	if err := s.Err(); err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("Result ", sumOfTwo*sumOfThree)

}

func charCount(s string) map[string]int {
	c := make(map[string]int)
	for i := 0; i < len(s); i++ {
		counter := strings.Count(s, string(s[i]))
		if counter == 2 || counter == 3 {
			c[string(s[i])] = counter
		}
	}
	return c
}

func mapper(c map[string]int) (int, int) {
	countOfTwo := 0
	countOfThree := 0
	for i, d := range c {
		if d == 2 {
			countOfTwo++
		}

		if d == 3 {
			countOfThree++
		}
		fmt.Println(i, " => ", d)
	}
	if countOfTwo > 1 {
		countOfTwo = 1
	}

	if countOfThree > 1 {
		countOfThree = 1
	}
	return countOfTwo, countOfThree
}
