package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		line_split := strings.Split(line, ",")
		range1 := strings.Split(line_split[0], "-")
		range2 := strings.Split(line_split[1], "-")
		range1_int := retype_arr(range1)
		range2_int := retype_arr(range2)
		if (range1_int[0] >= range2_int[0]) && (range1_int[1] <= range2_int[1]) {
			sum += 1
			continue
		}
		if (range2_int[0] >= range1_int[0]) && (range2_int[1] <= range1_int[1]) {
			sum += 1
			continue
		}

	}
	fmt.Println("sum: ", sum)
}

func part2() {
	var sum = 0
	var file, err = os.OpenFile("input.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		line_split := strings.Split(line, ",")
		range1 := strings.Split(line_split[0], "-")
		range2 := strings.Split(line_split[1], "-")
		range1_int := retype_arr(range1)
		range2_int := retype_arr(range2)
		if (range1_int[1] >= range2_int[0]) && (range1_int[0] <= range2_int[1]) {
			sum += 1
			continue
		}
		if (range2_int[1] >= range1_int[0]) && (range2_int[0] <= range1_int[1]) {
			sum += 1
			continue
		}
	}
	fmt.Println("sum: ", sum)
}

func retype_arr(arr []string) []int {
	var new_arr []int
	for _, val := range arr {
		new_val, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println(err.Error())
		}
		new_arr = append(new_arr, new_val)
	}
	return new_arr
}

func main() {
	part1()
	part2()
}
