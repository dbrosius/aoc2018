package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	allfreq := map[int]bool{0: true}
	thesum := 0
	var loopfreq *int
	for loopfreq == nil {
		file, err := os.Open("input.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			thetext := scanner.Text()
			theint, _ := strconv.Atoi(thetext)
			thesum += theint
			if allfreq[thesum] {
				loopfreq = &thesum
				break
			}
			allfreq[thesum] = true
		}
	}
	fmt.Println(*loopfreq)
}
