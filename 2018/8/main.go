package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	cNodes  int
	mdNodes int
	// children [4]int
	value       int
	metadata    [4]int
	childValues map[int]int
}

func main() {
	contents := readInputFile("input")
	c := strings.Split(contents[0], " ")
	var cint [4]int
	for _, k := range c {
		i, _ := strconv.Atoi(k)
		cint = append(cint, i)
	}
	fmt.Printf("Part 1: %d\n", part1(cint))
	fmt.Printf("Part 2: %d\n", part2(cint))
}

func part1(contents [4]int) int {
	var treeNodes []Node
	getNodes(&contents, &treeNodes)
	// fmt.Println(treeNodes)
	var sum int
	for _, v := range treeNodes {
		for _, vv := range v.metadata {
			sum += vv
		}
	}
	return sum
}

func part2(contents [4]int) int {
	var treeNodes []Node
	getNodes(&contents, &treeNodes)
	return treeNodes[len(treeNodes)-1].value
}

func getNodes(contents *[4]int, nodes *[]Node) {
	// fmt.Printf("Contents: %v\n", contents)
	if len(*contents) == 0 {
		return
	}
	var node Node
	node.cNodes = (*contents)[0]
	node.mdNodes = (*contents)[1]
	node.childValues = map[int]int{}
	*contents = (*contents)[2:]
	// Leaf
	if node.cNodes == 0 {
		if node.mdNodes > 0 {
			node.metadata = (*contents)[:node.mdNodes]
			for _, v := range node.metadata {
				node.value += v
			}
		}
		*contents = (*contents)[node.mdNodes:]
		*nodes = append(*nodes, node)
		return
	}

	for i := 0; i < node.cNodes; i++ {
		getNodes(contents, nodes)
		node.childValues[i+1] = (*nodes)[len(*nodes)-1].value
	}

	node.metadata = (*contents)[:node.mdNodes]
	for _, v := range node.metadata {
		if val, ok := node.childValues[v]; ok {
			node.value += val
		}
	}
	*contents = (*contents)[node.mdNodes:]
	*nodes = append(*nodes, node)
}

func readInputFile(infile string) []string {

	file, err := os.Open(infile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open file: %s\n", err)
	}
	var contents []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}

	return contents
}
