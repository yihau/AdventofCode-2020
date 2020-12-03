package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	_, file, _, _ := runtime.Caller(0)
	content, err := ioutil.ReadFile(filepath.Join(filepath.Dir(file), "./input.txt"))
	if err != nil {
		log.Fatalln("load input error:", err)
	}
	rawInput := strings.Split(string(content), "\n")

	fmt.Println("ans:",
		encounteredTree(rawInput, 1, 1)*
			encounteredTree(rawInput, 3, 1)*
			encounteredTree(rawInput, 5, 1)*
			encounteredTree(rawInput, 7, 1)*
			encounteredTree(rawInput, 1, 2))
}

func encounteredTree(road []string, eachStep, eachDown int) int {
	step := 1
	tree := 0
	for i := 0; i < len(road)-eachDown; i += eachDown {
		step += eachStep
		idx := step%len(road[i+eachDown]) - 1
		if idx == -1 {
			idx = len(road[i+eachDown]) - 1
		}
		if road[i+eachDown][idx] == '#' {
			tree++
		}
	}
	return tree
}
