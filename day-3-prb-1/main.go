package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type plot struct {
	id     int
	left   int
	top    int
	width  int
	height int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var rows []*plot
	s := bufio.NewScanner(f)
	for s.Scan() {
		var p plot
		_, err := fmt.Sscanf(s.Text(), "#%d @ %d,%d: %dx%d", &p.id, &p.left, &p.top, &p.width, &p.height)
		if err != nil {
			log.Fatal(err)
		}
		rows = append(rows, &p)
	}

	var a [1000][1000]string

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			a[i][j] = "."
		}
	}

	counter := 0
	for _, row := range rows {
		for h := row.top; h < row.top+row.height; h++ {
			for w := row.left; w < row.left+row.width; w++ {
				if a[h][w] == "." {
					a[h][w] = strconv.Itoa(row.id)

				} else if a[h][w] == "x" {
					continue
				} else {
					a[h][w] = "x"
					counter++
				}
			}
		}
	}

	fmt.Println(counter)

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
