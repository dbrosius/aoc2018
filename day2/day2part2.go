package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var all []string
	for scanner.Scan() {
		thetext := scanner.Text()
		all = append(all, thetext)
	}
	for idx, thestr := range all {
		idxsubslice := idx + 1
		subslice := all[idxsubslice:]
		output, ismatch := compareList(&subslice, &thestr)
		if ismatch {
			fmt.Println(output)
			break
		}
	}
}

func compareList(list *[]string, thevalue *string) (string, bool) {
	for _, str1 := range *list {
		output, ismatch := compareStrings(&str1, thevalue)
		if ismatch {
			return output, true
		}
	}
	return "", false
}

func compareStrings(str1 *string, str2 *string) (string, bool) {
	nonoverlap := 0
	output := ""
	for idx, therune1 := range *str1 {
		therune2 := []rune(*str2)[idx]
		if therune1 != therune2 {
			nonoverlap += 1
		} else {
			output = output + string(therune1)
		}
	}
	if nonoverlap == 1 {
		return output, true
	}
	return "", false
}
