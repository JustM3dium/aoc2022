package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var same = map[string]string{"A": "X", "B": "Y", "C": "Z"}
	var score = map[string]int64{"Y": 2, "X": 1, "Z": 3}
	var t string = "ABC"
	file, err := os.Open("resources/input.txt")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var finalScore int64
	for scanner.Scan() {
		splitted := strings.Split(scanner.Text(), string(' '))
		if splitted[1] == "X" {
			index := (strings.Index(t, splitted[0]) - 1) % 3
			if index == -1 {
				index = 2
			}
			finalScore += score[same[string(t[index])]]
		} else if splitted[1] == "Y" {
			finalScore += 3 + score[same[splitted[0]]]
		} else {
			index := (strings.Index(t, splitted[0]) + 1) % 3

			finalScore += score[same[string(t[index])]] + 6
		}
	}
	// First solution
	/*	if splitted[0] == rules[splitted[1]] {
			finalScore += score[splitted[1]]
		} else if splitted[1] == same[splitted[0]] {
			finalScore += 3 + score[splitted[1]]
		} else {
			finalScore += 6 + score[splitted[1]]
		}
	}*/

	fmt.Println(finalScore)

}
