package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
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
	var rucksackSave []string
	var lastVal string = ""
	for scanner.Scan() {
		tmpval := scanner.Text()
		if len(rucksackSave) < 3 {
			if lastVal != ""{
				rucksackSave = append(rucksackSave, lastVal)
				lastVal = ""
			}
			rucksackSave = append(rucksackSave, tmpval)
		} else if len(rucksackSave) == 3{
			lastVal = tmpval
			for i := 0; i < len(rucksackSave[0]); i++ {
				iVar := string(rucksackSave[0][i])
				if strings.Contains(rucksackSave[1], iVar) && strings.Contains(rucksackSave[2], iVar) &&
					!found{
					found = true
					finalScore += strings.Index(alphabet, string(rucksackSave[0][i])) + 1
					} else if found{
						break
				}
			}
		}
		if found {
			rucksackSave = rucksackSave[:0]
			found = false

		}
		//rucksackSave = rucksackSave[:0]
	}
	for i := 0; i < len(rucksackSave[0]); i++ {
		iVar := string(rucksackSave[0][i])
		if strings.Contains(rucksackSave[1], iVar) && strings.Contains(rucksackSave[2], iVar) &&
			!found {
			found = true
			finalScore += strings.Index(alphabet, string(rucksackSave[0][i])) + 1
		} else if found {
			break
		}
	}
	fmt.Println(finalScore)
	/*		scannerText := scanner.Text()
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
	*/
	//found = false

}