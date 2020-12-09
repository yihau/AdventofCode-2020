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
	rawInput = append(rawInput, "")

	m := map[byte]int{}
	people := 0
	total := 0
	for i := 0; i < len(rawInput); i++ {
		if rawInput[i] == "" {
			for _, times := range m {
				if times == people {
					total++
				}
			}
			people = 0
			m = map[byte]int{}
			continue
		}
		answered := map[byte]bool{}
		for j := 0; j < len(rawInput[i]); j++ {
			if !answered[rawInput[i][j]] {
				answered[rawInput[i][j]] = true
				m[rawInput[i][j]]++
			}
		}
		people++
	}
	fmt.Println("total:", total)
}
