package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	start := time.Now()

	var myInputNumbers []int
	for s.Scan() {
		d := s.Text()
		num, _ := strconv.Atoi(d)
		myInputNumbers = append(myInputNumbers, num)
	}
	count := make(map[int]int)
	_ = looper(myInputNumbers, count, 0)
	elapsed := time.Since(start)
	fmt.Println("Elapsed time is ", elapsed)
	//fmt.Println(result)
}

func looper(s []int, count map[int]int, lastItem int) map[int]int {
	count = insideLoop(s, count, lastItem)
	return count
}

func insideLoop(s []int, count map[int]int, lastItem int) map[int]int {
	var sum int
	var keys []int
	for i, x := range s {
		if i == 0 {
			//fmt.Println("The lastItem ", lastItem)
			sum += x + lastItem
		} else {
			sum += x
		}
		if i != 0 && count != nil {
			count[sum]++
		}
		keys = append(keys, sum)
		if count[sum] > 1 {
			fmt.Println("result is ", sum)
			break
		}
		//fmt.Println("The count of ", sum, " is ", count[sum])
	}

	hasTwos := false
	for _, d := range count {
		if d > 1 {
			hasTwos = true
			//fmt.Println("First time reaching", k, d)
			break
		}

	}
	if hasTwos == false {
		lastItem = keys[len(keys)-1]
		looper(s, count, lastItem)
	}
	return count
}
