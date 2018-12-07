package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	thesum := 0
	allfreq := map[int]bool{0: true}
	for scanner.Scan() {
		thetext := scanner.Text()
		theint, _ := strconv.Atoi(thetext)
		thesum += theint
		if allfreq[thesum] {
			fmt.Println(thesum)
		}
		allfreq[thesum] = true
	}
	fmt.Println(thesum)
}
