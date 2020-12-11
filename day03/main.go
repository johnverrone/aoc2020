package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("💩")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	trees := []int{0, 0, 0, 0, 0}
	xPos := []int{0, 0, 0, 0, 0}
	yPos := []int{0, 0, 0, 0, 0}
	slopes := make([][]int, 5)
	slopes[0] = []int{1, 1}
	slopes[1] = []int{3, 1}
	slopes[2] = []int{5, 1}
	slopes[3] = []int{7, 1}
	slopes[4] = []int{1, 2}

	for scanner.Scan() {
		line := scanner.Text()
		for i, slope := range slopes {
			if yPos[i]%slope[1] == 0 {
				if string(line[xPos[i]]) == "#" {
					trees[i]++
				}
				xPos[i] = (xPos[i] + slope[0]) % len(line)
			}
			yPos[i]++
		}
	}

	fmt.Printf("Trees encountered: %v\n", trees)
	ans := multiplyArray(trees...)
	fmt.Printf("Answer: %d\n", ans)
}

func multiplyArray(nums ...int) int {
	res := 1
	for _, num := range nums {
		res *= num
	}
	return res
}
