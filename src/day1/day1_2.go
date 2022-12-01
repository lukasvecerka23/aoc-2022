package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var count int
	list := []int{}
	var file, err = os.OpenFile("input.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line = scanner.Text()
		if len(line) == 0 {
			list = append(list, count)
			count = 0
			continue
		}
		var intLine, err = strconv.Atoi(line)
		if err != nil {
			fmt.Println(err.Error())
		}
		count = count + intLine
	}
	sort.Sort(sort.Reverse(sort.IntSlice(list)))
	fmt.Println("first: ", list[0])
	fmt.Println("second: ", list[1])
	fmt.Println("third: ", list[3])
	var sum = list[0] + list[1] + list[2]
	fmt.Println("sum: ", sum)
	if err := scanner.Err(); err != nil {
		fmt.Println(err.Error())
	}

}
