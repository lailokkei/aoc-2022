package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	rock     = 0
	paper    = 1
	scissors = 2
)

var movePoints = []int{1, 2, 3}

func outcome(self int, opponent int) int {
	if self == rock && opponent == scissors {
		return 6
	}
	if self == scissors && opponent == rock {
		return 0
	}
	if self > opponent {
		return 6
	}
	if self < opponent {
		return 0
	}
	return 3
}

func response(opponent int, outcome string) int {
	switch outcome {
	case "Y":
		return 3 + movePoints[opponent]
	case "X":
		if opponent == rock {
			return movePoints[scissors]
		}
		return movePoints[opponent-1]
	case "Z":
		score := 6
		if opponent == scissors {
			return score + movePoints[rock]
		}
		return score + movePoints[opponent+1]
	}
	return 0
}

func expect(got int, should int) {
	if got == should {
		fmt.Println("passed")
		return
	}
	fmt.Printf("got %d, was %d\n", got, should)
}

func main() {
	moveMap := map[string]int{
		"A": rock,
		"B": paper,
		"C": scissors,
		"X": rock,
		"Y": paper,
		"Z": scissors,
	}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var score int
	for scanner.Scan() {
		moves := strings.Split(scanner.Text(), " ")
		opponent := moveMap[moves[0]]
		outcome := moves[1]
		score += response(opponent, outcome)
	}
	fmt.Println(score)

}
