package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	m := make(map[int]bool)
	var rows []*plot
	s := bufio.NewScanner(f)
	for s.Scan() {
		var p plot
		_, err := fmt.Sscanf(s.Text(), "#%d @ %d,%d: %dx%d", &p.id, &p.left, &p.top, &p.width, &p.height)
		if err != nil {
			log.Fatal(err)
		}
		m[p.id] = false
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
		m[row.id] = false
		for h := row.top; h < row.top+row.height; h++ {
			for w := row.left; w < row.left+row.width; w++ {
				if a[h][w] == "." {
					a[h][w] = strconv.Itoa(row.id)

				} else if strings.Contains(a[h][w], "x") {

					var id int
					_, err := fmt.Sscanf(a[h][w], "%d-x", &id)
					if err != nil {
						log.Fatal(err)
					}

					m[id] = true
					m[row.id] = true
					continue
				} else {
					var id int
					_, err = fmt.Sscanf(a[h][w], "%d", &id)
					m[id] = true

					a[h][w] = strconv.Itoa(row.id) + "-x"

					m[row.id] = true
					counter++
				}
			}
		}

	}

	fmt.Println(counter)
	for k, r := range m {
		if !r {
			fmt.Println("The result is ", k)
			return
		}
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
