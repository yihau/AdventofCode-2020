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

	rawInput = append(rawInput, "")
	passport := map[string]string{}
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
			passport[kv[0]] = kv[1]
		}
	}
	fmt.Println("valid:", valid)
}

func emptyPassport(m map[string]string) {
	m["byr"] = ""
	m["iyr"] = ""
	m["eyr"] = ""
	m["hgt"] = ""
	m["hcl"] = ""
	m["ecl"] = ""
	m["pid"] = ""
	m["cid"] = ""
}

func isPassportValid(m map[string]string) bool {
	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	if m["byr"] == "" {
		return false
	}
	byr, err := strconv.Atoi(m["byr"])
	if err != nil {
		panic(err)
	}
	if byr < 1920 || byr > 2002 {
		return false
	}

	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	if m["iyr"] == "" {
		return false
	}
	iyr, err := strconv.Atoi(m["iyr"])
	if err != nil {
		panic(err)
	}
	if iyr < 2010 || iyr > 2020 {
		return false
	}

	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	if m["eyr"] == "" {
		return false
	}
	eyr, err := strconv.Atoi(m["eyr"])
	if err != nil {
		panic(err)
	}
	if eyr < 2020 || eyr > 2030 {
		return false
	}

	// hgt (Height) - a number followed by either cm or in:
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	if len(m["hgt"]) == 5 {
		if m["hgt"][3] != 'c' || m["hgt"][4] != 'm' {
			return false
		}
		hgt, err := strconv.Atoi(m["hgt"][:3])
		if err != nil {
			panic(err)
		}
		if hgt < 150 || hgt > 193 {
			return false
		}
	} else if len(m["hgt"]) == 4 {
		if m["hgt"][2] != 'i' || m["hgt"][3] != 'n' {
			return false
		}
		hgt, err := strconv.Atoi(m["hgt"][:2])
		if err != nil {
			panic(err)
		}
		if hgt < 59 || hgt > 76 {
			return false
		}
	} else {
		return false
	}

	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	if len(m["hcl"]) != 7 {
		return false
	}
	if m["hcl"][0] != '#' {
		return false
	}
	for i := 1; i < 7; i++ {
		ch := m["hcl"][i]
		if ch >= 'a' && ch <= 'f' {
			continue
		}
		if ch >= '0' && ch <= '9' {
			continue
		}
		return false
	}

	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	if m["ecl"] != "amb" &&
		m["ecl"] != "blu" &&
		m["ecl"] != "brn" &&
		m["ecl"] != "gry" &&
		m["ecl"] != "grn" &&
		m["ecl"] != "hzl" &&
		m["ecl"] != "oth" {
		return false
	}

	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	if len(m["pid"]) != 9 {
		return false
	}
	for i := 0; i < 9; i++ {
		if m["pid"][i] < '0' || m["pid"][i] > '9' {
			return false
		}
	}

	return true
}
