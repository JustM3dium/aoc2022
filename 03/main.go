package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main (){
    const alphabet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	file, err := os.Open("resources/input.txt")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
    //for scanner.Scan(){}
	var finalScore int = 0
	var found bool = false
	for scanner.Scan(){
		scannerText := scanner.Text()
		stringLen := len(scanner.Text())

		for i := 0 ; i <= (stringLen/2)-1; i++{
			iVal := scannerText[i]
			for j := stringLen/2; j < stringLen; j++{
				jVal := scannerText[j]
				if iVal == jVal && !found{
					finalScore += strings.Index(alphabet,string(iVal)) +1
					found = true
				}
			}
		}
		found = false

	}

	fmt.Println(finalScore)

}
