package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type column struct {
	crates []uint8
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

func part1() {
	var file, err = os.OpenFile("input.txt", os.O_RDONLY, 0644)
	var columns []column
	for i := 0; i < 9; i++ {
		columns = append(columns, column{[]uint8{}})
	}

	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		crate_num := 0
		line := scanner.Text()
		if len(line) == 0 || line[1] == '1' {
			continue
		}
		if !(strings.Contains(line, "move")) {
			backspc := 0
			line = strings.Replace(line, "[", "", -1)
			line = strings.Replace(line, "]", "", -1)
			for i := range line {
				if line[i] == ' ' {
					backspc++
				} else {
					columns[crate_num].crates = append(columns[crate_num].crates, line[i])
					crate_num++
					backspc = 0
					continue
				}
				if backspc == 4 {
					crate_num++
					backspc = 0
					continue
				}

			}
			continue
		}
		line = strings.Replace(line, "move ", "", -1)
		line = strings.Replace(line, "from ", "", -1)
		line = strings.Replace(line, "to ", "", -1)
		values := strings.Split(line, " ")
		values_int := retype_arr(values)
		for i := 0; i < values_int[0]; i++ {
			if len(columns[values_int[1]-1].crates) != 0 {
				columns[values_int[2]-1].crates = append([]uint8{columns[values_int[1]-1].crates[0]}, columns[values_int[2]-1].crates...)
				columns[values_int[1]-1].crates = columns[values_int[1]-1].crates[1:]
			}
		}
	}
	fmt.Println(columns)
}

func part2() {
	var file, err = os.OpenFile("input.txt", os.O_RDONLY, 0644)
	var columns []column
	for i := 0; i < 9; i++ {
		columns = append(columns, column{[]uint8{}})
	}

	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		crate_num := 0
		line := scanner.Text()
		if len(line) == 0 || line[1] == '1' {
			continue
		}
		if !(strings.Contains(line, "move")) {
			backspc := 0
			line = strings.Replace(line, "[", "", -1)
			line = strings.Replace(line, "]", "", -1)
			for i := range line {
				if line[i] == ' ' {
					backspc++
				} else {
					columns[crate_num].crates = append(columns[crate_num].crates, line[i])
					crate_num++
					backspc = 0
					continue
				}
				if backspc == 4 {
					crate_num++
					backspc = 0
					continue
				}

			}
			continue
		}
		line = strings.Replace(line, "move ", "", -1)
		line = strings.Replace(line, "from ", "", -1)
		line = strings.Replace(line, "to ", "", -1)
		values := strings.Split(line, " ")
		values_int := retype_arr(values)
		if len(columns[values_int[1]-1].crates) != 0 {
			tmp_arr := make([]uint8, values_int[0])
			copy(tmp_arr, columns[values_int[1]-1].crates[:values_int[0]])
			columns[values_int[2]-1].crates = append(tmp_arr, columns[values_int[2]-1].crates...)
			columns[values_int[1]-1].crates = columns[values_int[1]-1].crates[values_int[0]:]
		}
	}
	fmt.Println(columns)
}

func main() {
	part1()
	part2()
}
