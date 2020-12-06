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

	m := map[int]bool{}
	for _, input := range rawInput {
		m[GetNum('B', 'F', input[:7], 128)*8+GetNum('R', 'L', input[7:], 8)] = true
	}

	for row := 1; row < 127; row++ {
		for col := 0; col < 8; col++ {
			if !m[row*8+col] && m[row*8+col+1] && m[row*8+col-1] {
				fmt.Println("find", row*8+col)
			}
		}
	}
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
