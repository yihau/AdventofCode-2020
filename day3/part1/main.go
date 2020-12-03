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

	step := 1
	tree := 0
	for i := 0; i < len(rawInput)-1; i++ {
		step += 3
		idx := step%len(rawInput[i+1]) - 1
		if idx == -1 {
			idx = len(rawInput[i+1]) - 1
		}
		if rawInput[i+1][idx] == '#' {
			tree++
		}
	}
	fmt.Println("Tree:", tree)
}
