package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	_, file, _, _ := runtime.Caller(0)
	content, err := ioutil.ReadFile(filepath.Join(filepath.Dir(file), "./input.txt"))
	if err != nil {
		log.Fatalln("load input error:", err)
	}
	rawInput := strings.Split(string(content), "\n")
	valid := 0
	for i := 0; i < len(rawInput); i++ {
		input := strings.Split(rawInput[i], " ")
		appearRange := strings.Split(input[0], "-")
		pos1, err := strconv.Atoi(appearRange[0])
		if err != nil {
			log.Fatalln("input not number", err)
		}
		pos2, err := strconv.Atoi(appearRange[1])
		if err != nil {
			log.Fatalln("input not number", err)
		}
		target := input[1][0]
		if (input[2][pos1-1] == target && input[2][pos2-1] != target) || (input[2][pos1-1] != target && input[2][pos2-1] == target) {
			valid++
		}
	}
	fmt.Println("valid:", valid)
	fmt.Println("done")
}
