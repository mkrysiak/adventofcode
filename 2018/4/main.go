package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var startRE = regexp.MustCompile(`\[\d+\-(\d\d\-\d\d)\s(\d+):(\d+)\]\sGuard\s#(\d+)\sbegins\sshift`)
var asleepRE = regexp.MustCompile(`\[\d+\-(\d\d\-\d\d)\s(\d+):(\d+)\]\sfalls\sasleep`)
var awakeRE = regexp.MustCompile(`\[\d+\-(\d\d\-\d\d)\s(\d+):(\d+)\]\swakes\sup`)

func main() {
	contents := readInputFile("input")
	sort.Strings(contents)
	fmt.Printf("Part 1: %d\n", part1(contents))
	fmt.Printf("Part 2: %d\n", part2(contents))
}

func part1(contents []string) int {
	// map[guardId] = total minutes asleep
	asleepGuard := map[string]int{}
	// map[guardId-minute] = count of times asleep during this minute
	asleepMap := map[string]int{}
	var fallsAsleep, wakesUp int
	var id string
	for _, v := range contents {
		if start := startRE.FindStringSubmatch(v); start != nil {

			id = start[4]
			if _, ok := asleepGuard[id]; !ok {
				asleepGuard[id] = 0
			}

		} else if asleep := asleepRE.FindStringSubmatch(v); asleep != nil {
			fallsAsleep, _ = strconv.Atoi(asleep[3])

		} else if awake := awakeRE.FindStringSubmatch(v); awake != nil {
			wakesUp, _ = strconv.Atoi(awake[3])

			asleepGuard[id] += (wakesUp - fallsAsleep)
			for i := fallsAsleep; i < wakesUp; i++ {
				si := strconv.Itoa(i)
				asleepMap[id+"-"+si]++
			}
		}
	}

	sleepiestGuard := ""
	mostSleep := 0
	for k, v := range asleepGuard {
		if v > mostSleep {
			sleepiestGuard = k
			mostSleep = v
		}
	}

	sleepiestMinute := 0
	sleepiestCount := 0
	for i := 0; i < 59; i++ {
		uid := strconv.Itoa(i)
		key := sleepiestGuard + "-" + uid
		if sleepiestCount < asleepMap[key] {
			sleepiestCount = asleepMap[key]
			sleepiestMinute = i
		}
	}
	s, _ := strconv.Atoi(sleepiestGuard)

	return s * sleepiestMinute
}

func part2(contents []string) int {
	asleepGuard := map[string]int{}
	asleepMap := map[string]int{}
	var fallsAsleep, wakesUp int
	var id string
	for _, v := range contents {
		if start := startRE.FindStringSubmatch(v); start != nil {

			id = start[4]
			if _, ok := asleepGuard[id]; !ok {
				asleepGuard[id] = 0
			}

		} else if asleep := asleepRE.FindStringSubmatch(v); asleep != nil {
			fallsAsleep, _ = strconv.Atoi(asleep[3])

		} else if awake := awakeRE.FindStringSubmatch(v); awake != nil {
			wakesUp, _ = strconv.Atoi(awake[3])

			asleepGuard[id] += (wakesUp - fallsAsleep)
			for i := fallsAsleep; i < wakesUp; i++ {
				si := strconv.Itoa(i)
				asleepMap[id+"-"+si]++
			}
		}
	}

	sleepiestMinute := 0
	sleepiestGuard := ""
	for k, v := range asleepMap {
		if v > sleepiestMinute {
			sleepiestMinute = v
			sleepiestGuard = k
		}
	}
	guardIdAndMinute := strings.Split(sleepiestGuard, "-")
	guardId, _ := strconv.Atoi(guardIdAndMinute[0])
	minute, _ := strconv.Atoi(guardIdAndMinute[1])

	return guardId * minute
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
