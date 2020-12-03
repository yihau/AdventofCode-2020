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
	input := make([]int, 0, len(rawInput))
	for i := 0; i < len(rawInput); i++ {
		num, err := strconv.Atoi(rawInput[i])
		if err != nil {
			log.Fatalln("input not a number:", err)
		}
		input = append(input, num)
	}

	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input)-1; j++ {
			for k := j + 1; k < len(input)-2; k++ {
				if 2020-input[i]-input[j]-input[k] == 0 {
					fmt.Println("find:", input[i]*input[j]*input[k])
				}
			}
		}
	}
	fmt.Println("done")
}
