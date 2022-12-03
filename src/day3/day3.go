package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func part1() {
	var sum = 0
	var file, err = os.OpenFile("input.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var line = scanner.Text()
		var half1 = line[:len(line)/2]
		var half2 = line[len(line)/2:]

		var char = char_contains1(half1, half2)

		if unicode.IsLower(char) {
			sum += int(char) - 96
		} else {
			sum += int(char) - 38
		}
	}
	fmt.Println("sum: ", sum)
}

func char_contains1(line1 string, line2 string) int32 {
	for _, char := range line1 {
		if strings.ContainsRune(line2, char) {
			return char
		}
	}
	return 0
}

func char_contains2(line1 string, line2 string) []int32 {
	var ret = []int32{}
	for _, char := range line1 {
		if strings.ContainsRune(line2, char) {
			ret = append(ret, char)
		}
	}
	return ret
}

func equal_char(arr1 []int32, arr2 []int32) int32 {
	for _, v := range arr1 {
		for _, val := range arr2 {
			if v == val {
				return v
			}
		}
	}
	return 0
}

func part2() {
	var sum = 0
	var file, err = os.OpenFile("input.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var pre_line string
	var line_count = 0
	var letters1 []int32
	for scanner.Scan() {
		line_count++
		var line = scanner.Text()
		if line_count == 2 {
			letters1 = char_contains2(line, pre_line)
		} else if line_count == 3 {
			var letters2 = char_contains2(line, pre_line)
			var letter = equal_char(letters1, letters2)
			if unicode.IsLower(letter) {
				sum += int(letter) - 96
			} else {
				sum += int(letter) - 38
			}
			line_count = 0
		}
		pre_line = line
	}
	fmt.Println("sum: ", sum)
}

func main() {
	part1()
	part2()
}
