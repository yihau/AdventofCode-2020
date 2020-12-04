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
	passport := map[string]bool{}
	emptyPassport(passport)
	valid := 0
	for i := 0; i < len(rawInput); i++ {
		if rawInput[i] == "" {
			if isPassportValid(passport) {
				valid++
			}
			emptyPassport(passport)
			continue
		}
		for _, input := range strings.Split(rawInput[i], " ") {
			kv := strings.Split(string(input), ":")
			passport[kv[0]] = true
		}
	}
	fmt.Println("valid:", valid)
}

func emptyPassport(m map[string]bool) {
	m["byr"] = false
	m["iyr"] = false
	m["eyr"] = false
	m["hgt"] = false
	m["hcl"] = false
	m["ecl"] = false
	m["pid"] = false
	m["cid"] = false
}

func isPassportValid(m map[string]bool) bool {
	return m["byr"] && m["iyr"] && m["eyr"] && m["hgt"] && m["hcl"] && m["ecl"] && m["pid"]
}
