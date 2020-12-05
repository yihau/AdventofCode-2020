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

	result := 0
	for _, input := range rawInput {
		result = max(result, GetNum('B', 'F', input[:7], 128)*8+GetNum('R', 'L', input[7:], 8))
	}
	fmt.Println(result)
}

func GetNum(rightSign, leftSign byte, str string, total int) int {
	left, right := 0, total-1
	for i := 0; i < len(str); i++ {
		if str[i] == rightSign {
			left += ((right - left) + 1) / 2
		} else {
			right -= ((right - left) + 1) / 2
		}
	}
	return left
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
