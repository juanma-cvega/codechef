package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	minCases = 1
	maxCases = 1000
	minSets = 1
	maxSets = 10 * 5
	repsInSet = 15
)

var reader *bufio.Reader

func main() {
    reader = bufio.NewReader(os.Stdin)

	cases := readCases()
	sets := readSets(cases)
	res := calculateRepsFrom(sets)
	for _, caseRes := range res {
		writeResultFrom(strconv.Itoa(caseRes))
	}
}

func calculateRepsFrom(sets []int) []int{
	var res []int
	for _, set := range sets {
		res = append(res, set * repsInSet)
	}
	return res
}

func readCases() int {
    l := readNewLine()
    v, err := strconv.Atoi(l)
    if err != nil {
        panic(fmt.Errorf("Error while transforming input into number: %s", err))
    }
    if v < minCases && v > maxCases {
        panic(fmt.Errorf("Number of cases is not within range (1 - 100)"))
    }
    return v
}

func readSets(cases int) []int {
	var sets []int
	for i := 0; i < cases; i++ {
		l := readNewLine()
		set, err := strconv.Atoi(l)
		if err != nil {
			panic(fmt.Errorf("unable to get value: %s", err))
		}
		sets = append(sets, set)
	}
	return sets
}

func readNewLine() string {
    text, err := reader.ReadString('\n')
    if err != nil {
        panic(fmt.Errorf("Error reading from input: %s", err))
    }
    return strings.TrimSuffix(text, "\n")
}

func writeResultFrom(s string) {
    fmt.Println(s)
}