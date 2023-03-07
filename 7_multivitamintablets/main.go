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
	maxCases = 2000
	tabletsPerDay = 3
	minDays = 1
	maxDays = 100
	minTablets = 0
	maxTablets = 1000
)

var reader bufio.Reader = *bufio.NewReader(os.Stdin)

func main() {
	nCases := readCases()
	for i := 0;i < nCases;i++ {
		input := readInput()
		writeResultFrom(enoughTablets(input))
	}
}

func enoughTablets(input input) string{
	if input.days * tabletsPerDay <= input.tablets {
		return "YES"
	}
	return "NO"
}

type input struct {
	days, tablets int
}

func newFromArray(vv []string) input {
	if len(vv) != 2 {
		panic("number of values for the case invalid")
	}
	return new(vv[0], vv[1])
}

func new(days, tablets string) input {
	return input{
		days: convert(days),
		tablets: convert(tablets),
	}
}

func convert(value string) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		panic(fmt.Errorf("error converting to int: %s", err))
	}
	return i
}

func readInput() input {
	vv := strings.Split(readNewLine(), " ")
	return newFromArray(vv)
}

func readCases() int {
	l := readNewLine()
	v, err := strconv.Atoi(l)
	if err != nil {
		panic(fmt.Errorf("error while transforming input into number: %s", err))
	}
	if v < minCases && v > maxCases {
		panic(fmt.Errorf("number of cases is not within range (1 - 100)"))
	}
	return v
}

func readNewLine() string {
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(fmt.Errorf("error reading from input: %s", err))
	}
	return strings.TrimSuffix(text, "\n")
}

func writeResultFrom(result string){
	fmt.Println(result)
}