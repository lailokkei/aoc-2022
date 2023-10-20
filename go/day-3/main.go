package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

func priority(char rune) (int, error) {
	value := int(char)
	if value >= 65 && value <= 90 {
		return value - 38, nil
	}
	if value >= 97 && value <= 122 {
		return value - 96, nil
	}
	return value, errors.New("not a valid letter")
}

func compare(x string, y string) int {
	for _, i := range x {
		for _, j := range y {
			if i == j {
				idk, _ := priority(rune(i))
				return idk
			}
		}
	}
	return 35498
}

func main() {
	fmt.Println(priority('a'))
	fmt.Println(priority('A'))

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		rucksack := scanner.Text()
		comp1 := rucksack[:len(rucksack)/2]
		comp2 := rucksack[len(rucksack)/2:]
		total += compare(comp1, comp2)
	}
	fmt.Println(total)
}
