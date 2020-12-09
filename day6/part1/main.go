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

	m := map[byte]bool{}
	total := 0
	for i := 0; i < len(rawInput); i++ {
		if rawInput[i] == "" {
			total += len(m)
			m = map[byte]bool{}
			continue
		}
		for j := 0; j < len(rawInput[i]); j++ {
			m[rawInput[i][j]] = true
		}
	}
	fmt.Println("total:", total)
}
