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
	input := strings.Split(string(content), "\n")
	m := map[int]bool{}
	for i := 0; i < len(input); i++ {
		num, err := strconv.Atoi(input[i])
		if err != nil {
			log.Fatalln("input not a number:", err)
		}
		if m[2020-num] {
			fmt.Println("find:", num*(2020-num))
			break
		}
		m[num] = true
	}
	fmt.Println("done")
}
