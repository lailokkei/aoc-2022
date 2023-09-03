package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func max(a int, b int) int {
	if b > a {
		return b
	}
	return a
}

func total(items []string) int {
	total := 0
	for i := 0; i < len(items); i++ {
		value, err := strconv.Atoi(items[i])
		if err != nil {
			return total
		}
		total += value
	}
	return total
}

func highest(lines string) int {
	highest := 0

	elves := strings.Split(lines, "\n\n")
	for i := 0; i < len(elves); i++ {
		items := strings.Split(elves[i], "\n")
		highest = max(highest, total(items))
	}

	return highest
}

func updateRank(top3 [3]int, newValue int) [3]int {
	for i, val := range top3 {
		if newValue > val {
			for j := 2; j > i; j-- {
				top3[j] = top3[j-1]
			}
			top3[i] = newValue
			return top3
		}
	}
	return top3
}

func highest3(lines string) int {
	var top3 [3]int

	elves := strings.Split(lines, "\n\n")
	for i := 0; i < len(elves); i++ {
		items := strings.Split(elves[i], "\n")
		top3 = updateRank(top3, total(items))
	}

	var total int
	for i := 0; i < len(top3); i++ {
		total += top3[i]
	}
	log.Println(top3)
	return total
}

func main() {

	file, err := os.ReadFile("calories.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := string(file)

	fmt.Printf("highest : %d\n", highest(lines))
	fmt.Printf("top 3 total : %d\n", highest3(lines))

	// var rank [3]int
	// fmt.Println(rank)
	// rank = updateRank(rank, 34)
	// fmt.Println(rank)
	// rank = updateRank(rank, 40)
	// fmt.Println(rank)
	// rank = updateRank(rank, 20)
	// fmt.Println(rank)
	// rank = updateRank(rank, 123)
	// fmt.Println(rank)
}
