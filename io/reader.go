package io

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader bufio.Reader = *bufio.NewReader(os.Stdin)

type Input struct {
	ValueOne, ValueTwo int
}

func NewFromArray(vv []string) Input {
	if len(vv) != 2 {
		panic("number of values for the case invalid")
	}
	return New(vv[0], vv[1])
}

func New(value1, value2 string) Input {
	return Input{
		ValueOne: convert(value1),
		ValueTwo: convert(value2),
	}
}

func convert(value string) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		panic(fmt.Errorf("error converting to int: %s", err))
	}
	return i
}

func ReadInput() Input {
	vv := strings.Split(ReadNewLine(), " ")
	return NewFromArray(vv)
}

func ReadCases(minCases, maxCases int) int {
	l := ReadNewLine()
	v, err := strconv.Atoi(l)
	if err != nil {
		panic(fmt.Errorf("error while transforming input into number: %s", err))
	}
	if v < minCases && v > maxCases {
		panic(fmt.Errorf("number of cases is not within range (1 - 100)"))
	}
	return v
}

func ReadNewLine() string {
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(fmt.Errorf("error reading from input: %s", err))
	}
	return strings.TrimSuffix(text, "\n")
}

func WriteResultFrom(result string){
	fmt.Println(result)
}