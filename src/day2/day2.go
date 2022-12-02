package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var combinations = map[string]int{"A": 1, "B": 2, "C": 3, "X": 1, "Y": 2, "Z": 3}
	var score = 0
	// chapter 1
	var file, err = os.OpenFile("example", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line = scanner.Text()
		game := strings.Fields(line)
		fmt.Println(game)
		var combination_1 = combinations[game[0]]
		var combination_2 = combinations[game[1]]
		if combination_1 == combination_2 {
			score += combination_2 + 3
			continue
		} else if combination_1 == 1 && combination_2 == 3 {
			score += combination_2 + 0
			continue
		} else if combination_1 == 2 && combination_2 == 1 {
			score += combination_2 + 0
			continue
		} else if combination_1 == 3 && combination_2 == 2 {
			score += combination_2 + 0
			continue
		}
		score += combination_2 + 6
	}
	fmt.Println("total score: ", score)

	// chapter 2
	var file2, err2 = os.OpenFile("input.txt", os.O_RDONLY, 0644)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	defer file2.Close()

	scanner2 := bufio.NewScanner(file2)
	score = 0
	for scanner2.Scan() {
		var line = scanner2.Text()
		game := strings.Fields(line)
		fmt.Println(game)
		var combination_1 = combinations[game[0]]
		var result = combinations[game[1]]

		if combination_1 == 1 {
			switch result {
			case 1:
				score += 3 + 0
				break
			case 2:
				score += 1 + 3
				break
			case 3:
				score += 2 + 6
				break
			}
			continue
		} else if combination_1 == 2 {
			switch result {
			case 1:
				score += 1 + 0
				break
			case 2:
				score += 2 + 3
				break
			case 3:
				score += 3 + 6
			}
			continue
		} else if combination_1 == 3 {
			switch result {
			case 1:
				score += 2 + 0
				break
			case 2:
				score += 3 + 3
				break
			case 3:
				score += 1 + 6
				break
			}
			continue
		}
	}
	fmt.Println("total score 2: ", score)
}
