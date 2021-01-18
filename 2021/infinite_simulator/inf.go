package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

var RE = regexp.MustCompile("Puzzle \\d* \\(([^)*])\\)\t(.*)$")

func main() {
	puzzles := map[string][]string{}
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	rdr := bufio.NewReader(f)
	for {
		line, err := rdr.ReadString('\n')
		
	}
}
