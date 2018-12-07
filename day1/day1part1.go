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
	for scanner.Scan() {
		thetext := scanner.Text()
		theint, _ := strconv.Atoi(thetext)
		thesum += theint
	}
	fmt.Println(thesum)
}
