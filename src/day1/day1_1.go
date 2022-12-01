package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var count int
	var max = 0
	var file, err = os.OpenFile("input.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line = scanner.Text()
		if len(line) == 0 {
			if count > max {
				max = count
			}
			count = 0
			continue
		}
		var intLine, err = strconv.Atoi(line)
		if err != nil {
			fmt.Println(err.Error())
		}
		count = count + intLine
	}
	fmt.Println("max: ", max)
	if err := scanner.Err(); err != nil {
		fmt.Println(err.Error())
	}

}
