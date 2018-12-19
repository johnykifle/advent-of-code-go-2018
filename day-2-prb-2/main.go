package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var n []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		n = append(n, s.Text())
	}
	start := time.Now()

	getWords(n)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func getWords(n []string) {
	for _, word1 := range n {
		for _, word2 := range n {
			if word1 != word2 {
				length := len(word1)
				maxNumberOfTrue := length - 1
				var common []rune
				for i, character := range word2 {
					currenctCharacter1 := word1[i : i+1]
					result := currenctCharacter1 == string(character)
					if result {
						maxNumberOfTrue--
						common = append(common, character)
					}
					if maxNumberOfTrue == 0 {
						fmt.Println(string(common))
						return
					}
				}
			}
		}
	}
	return
}
