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
	maxCases = 100
	minInput = 1
	maxInput = 100
)

var reader *bufio.Reader

type input struct {
	earned int
	taxThreshold int
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	nCases := readCases()
	inputs := readInputs(nCases)
	for _, input := range inputs {
		writeResultFrom(calculateTaxFor(input))
	}
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

func readInputs(cases int) []input {
	var inputs []input
	for i := 0; i < cases; i++ {
		vv := strings.Split(readNewLine(), " ")
		if len(vv) != 2 {
			panic(fmt.Errorf("wrong number of inputs for case %d, got %d, wanted 2", i, len(vv)))
		}
		v1, err := strconv.Atoi(vv[0])
		if err != nil {
			panic(fmt.Errorf("unable to transform input: %s", err))
		}
		v2, err := strconv.Atoi(vv[1])
		if err != nil {
			panic(fmt.Errorf("unable to transform input: %s", err))
		}
		inputs = append(inputs, input{
			earned: v1,
			taxThreshold: v2,
		})
	}
	return inputs
}

func calculateTaxFor(i input) int{
	return i.earned - i.taxThreshold
}

func readNewLine() string {
    text, err := reader.ReadString('\n')
    if err != nil {
        panic(fmt.Errorf("Error reading from input: %s", err))
    }
    return strings.TrimSuffix(text, "\n")
}

func writeResultFrom(v int) {
    fmt.Println(strconv.Itoa(v))
}