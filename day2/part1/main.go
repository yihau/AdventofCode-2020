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
		min, err := strconv.Atoi(appearRange[0])
		if err != nil {
			log.Fatalln("input not number", err)
		}
		max, err := strconv.Atoi(appearRange[1])
		if err != nil {
			log.Fatalln("input not number", err)
		}
		target := input[1][0]
		appear := 0
		for j := 0; j < len(input[2]); j++ {
			if input[2][j] == target {
				appear++
			}
		}
		if appear >= min && appear <= max {
			valid++
		}
	}
	fmt.Println("valid:", valid)
	fmt.Println("done")
}
