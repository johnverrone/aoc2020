package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type policy func(int, int, string, string) bool

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("💩")
	}

	pws := strings.Split(string(data), "\n")

	// Part 1
	countPolicy := func(low int, high int, char string, password string) bool {
		count := strings.Count(password, char)
		return count >= low && count <= high
	}

	// Part 2
	positionPolicy := func(low int, high int, char string, password string) bool {
		return (string(password[low-1]) == char) != (string(password[high-1]) == char)
	}

	validCountPasswords := 0
	validPositionPasswords := 0
	for _, pw := range pws {
		if valid(pw, countPolicy) {
			validCountPasswords++
		}
		if valid(pw, positionPolicy) {
			validPositionPasswords++
		}
	}

	fmt.Printf("Valid passwords pt1: %d\n", validCountPasswords)
	fmt.Printf("Valid passwords pt2: %d\n", validPositionPasswords)
}

func valid(pw string, p policy) bool {
	r := regexp.MustCompile(`(\d*)-(\d*) (\D): (.*)`)
	match := r.FindStringSubmatch(pw)
	low, _ := strconv.Atoi(match[1])
	high, _ := strconv.Atoi(match[2])
	letter := match[3]
	pass := match[4]

	return p(low, high, letter, pass)
}
