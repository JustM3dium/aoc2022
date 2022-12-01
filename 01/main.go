package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	file, err := os.Open("resources/input.txt")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var calories int64
	var caloriesSlice []int64
	for scanner.Scan() {
		currentNumber, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			caloriesSlice = append(caloriesSlice, calories)
			calories = 0
		} else {
			calories += currentNumber
		}
	}

	caloriesSlice = append(caloriesSlice, calories)

	fErr := file.Close()
	if fErr != nil {
		return
	}
	sort.Slice(caloriesSlice, func(p, q int) bool {
		return caloriesSlice[p] > caloriesSlice[q]
	})

	first := caloriesSlice[0]
	second := caloriesSlice[1]
	third := caloriesSlice[2]

	//fmt.Println(caloriesSlice)
	fmt.Println("Most calories: ", first)
	//fmt.Println(first, second, third)
	fmt.Println("Sum of three most: ", first+second+third)
}
