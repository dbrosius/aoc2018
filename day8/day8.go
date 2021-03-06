package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type treenode struct {
	nodeCount int
	metaCount int
	metadata  []int
	nodes     []treenode
}

func parseNodes(intData []int, thesum int) (treenode, []int, int) {
	nodeCount := intData[0]
	metaCount := intData[1]
	var nodes []treenode
	remaining := intData[2:]
	for idx := 0; idx < nodeCount; idx++ {
		thenode, rem, newsum := parseNodes(remaining, thesum)
		thesum = newsum
		remaining = rem
		nodes = append(nodes, thenode)
	}
	metadata := remaining[0:metaCount]
	for _, theval := range metadata {
		thesum += theval
	}
	remaining = remaining[(0 + metaCount):]
	return treenode{nodeCount: nodeCount, metaCount: metaCount, metadata: metadata, nodes: nodes}, remaining, thesum
}

func sum(intData []int) int {
	thesum := 0
	for _, theval := range intData {
		thesum += theval
	}
	return thesum
}

func nodevalue(rootnode treenode) int {
	fmt.Println(rootnode.nodeCount)
	if rootnode.nodeCount == 0 {
		return sum(rootnode.metadata)
	}
	thesum := 0
	for _, nodeidx := range rootnode.metadata {
		if nodeidx <= rootnode.nodeCount && nodeidx > 0 {
			thesum += nodevalue(rootnode.nodes[nodeidx-1])
			// fmt.Println(thesum)
		}
	}
	return thesum
}

func main() {
	filebytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	filestr := string(filebytes)
	filearr := strings.Split(filestr, " ")
	var fileints []int
	for _, fileintstr := range filearr {
		fileint, _ := strconv.Atoi(fileintstr)
		fileints = append(fileints, fileint)
	}
	rootnode, _, thesum := parseNodes(fileints, 0)
	rootval := nodevalue(rootnode)
	fmt.Printf("The sum of the metadata is %d\n", thesum)
	fmt.Printf("The value of the root node is %d\n", rootval)
}
