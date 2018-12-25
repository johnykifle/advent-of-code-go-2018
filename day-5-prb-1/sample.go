package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	f, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var data []string
	input := string(f)
	for _, s := range input {
		data = append(data, string(s))
	}

	//var result []string
	for i, k := range data {
		if strings.EqualFold(k, data[i+1]) && (k != data[i+1]) {
			fmt.Println("dd")
		}
	}

}
