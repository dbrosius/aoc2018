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
	countmap := make(map[int]int)
	for scanner.Scan() {
		thetext := scanner.Text()
		thisline := make(map[rune]int)
		for _, therune := range thetext {
			thisline[therune] += 1
		}
		thiscount := make(map[int]int)
		for _, v := range thisline {
			thiscount[v] = 1
		}
		for k, v := range thiscount {
			countmap[k] += v
		}
	}
	fmt.Println(countmap[2]*countmap[3])
}
