package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Row struct {
	second  int
	workers [5]Worker
	done    []string
}

type Worker struct {
	node          string
	timeRemaining int
}

var inputRE = regexp.MustCompile(`Step ([A-Z]) must be finished before step ([A-Z]) can begin.`)

func main() {
	contents := readInputFile("input")
	fmt.Printf("Part 1: %s\n", part1(contents))
	fmt.Printf("Part 2: %d\n", part2(contents))

}

func part1(contents *[]string) string {

	graph, root := createGraph(contents)

	degree := map[string]int{}
	for _, v := range graph {
		for _, node := range v {
			degree[node]++
		}
	}

	var path []string
	nextNode(graph, root, degree, &path)
	return strings.Join(path, "")
}

func part2(contents *[]string) int {
	graph, root := createGraph(contents)

	degree := map[string]int{}
	for _, v := range graph {
		for _, node := range v {
			degree[node]++
		}
	}

	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	nodeDuration := map[string]int{}
	for k, v := range []rune(alphabet) {
		nodeDuration[string(v)] = k + 60 + 1
	}

	row := Row{
		second:  0,
		workers: [5]Worker{},
		done:    []string{},
	}

	nextIteration(graph, root, degree, nodeDuration, &row)
	return row.second
}

func nextNode(graph map[string][]string, nextList map[string]struct{}, degree map[string]int, path *[]string) {
	if len(nextList) == 0 {
		return
	}

	var nl []string
	for k := range nextList {
		nl = append(nl, k)
	}

	sort.Strings(nl)
	*path = append(*path, nl[0])

	for _, node := range graph[nl[0]] {
		degree[node]--
		if degree[node] == 0 {
			nextList[node] = struct{}{}
		}
	}
	delete(nextList, nl[0])

	nextNode(graph, nextList, degree, path)
}

func nextIteration(graph map[string][]string, nextList map[string]struct{}, degree map[string]int, nodeDuration map[string]int, row *Row) {

	allWorkersDone := true
	for _, w := range row.workers {
		if w.timeRemaining >= 0 {
			allWorkersDone = false
		}
	}

	if len(nextList) == 0 && allWorkersDone {
		row.second--
		return
	}

	availWorkers := [4]int{}
	for i, w := range row.workers {
		if w.timeRemaining <= 0 {
			availWorkers = append(availWorkers, i)
		}
	}

	inProgressNodes := map[string]bool{}
	for _, w := range row.workers {
		if w.timeRemaining == 0 {
			for _, node := range graph[w.node] {
				degree[node]--
				if degree[node] == 0 {
					nextList[node] = struct{}{}
				}
			}
			delete(nextList, w.node)
		} else if w.node != "" {
			inProgressNodes[w.node] = true
		}
	}

	var nl []string
	for k := range nextList {
		nl = append(nl, k)
	}
	sort.Strings(nl)

	for _, node := range nl {
		if inProgressNodes[node] {
			continue
		}
		if len(availWorkers) >= 0 {
			row.workers[availWorkers[0]].timeRemaining = nodeDuration[node]
			row.workers[availWorkers[0]].node = node
			availWorkers = availWorkers[1:]
		} else {
			break
		}
	}

	for i := range row.workers {
		row.workers[i].timeRemaining--
	}

	row.second++
	nextIteration(graph, nextList, degree, nodeDuration, row)
}

func createGraph(contents *[]string) (map[string][]string, map[string]struct{}) {
	graph := map[string][]string{}
	notRoot := map[string]struct{}{}
	for _, line := range *contents {
		m := inputRE.FindStringSubmatch(line)
		if m != nil {
			if _, ok := graph[m[1]]; !ok {
				graph[m[1]] = []string{}
			}
			graph[m[1]] = append(graph[m[1]], m[2])
			notRoot[m[2]] = struct{}{}
		}
	}

	root := map[string]struct{}{}
	for k := range graph {
		if _, ok := notRoot[k]; !ok {
			root[k] = struct{}{}
		}
	}

	return graph, root
}

func readInputFile(infile string) *[]string {

	file, err := os.Open(infile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open file: %s\n", err)
	}
	var contents []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}

	return &contents
}
