package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

const(
	minCases = 1
	maxCases = 1000
	minAge = 20
	maxAge = 40
	minInputAge = 10
	maxInputAge = 50
)

func main() {
	reader = bufio.NewReader(os.Stdin)
	nCases := readNumberOfCases()
	for i := 0; i < nCases; i++ {
		writeResult(isValid(readCase()))
	}
}

func writeResult(res string) {
	fmt.Println(res)
}

func isValid(x, y, a int) string{
	res := a >= x && a < y
	if res {
		return "YES"
	}
	return "NO"
}

func readCase() (x, y, a int){
	currentCase := readNewLine()
	values := strings.Split(currentCase, " ")
	if len(values) != 3 {
		panic("invalid case values")
	}
	caseValues := make([]int, 3)
	var err error
	for i, v := range values {
		caseValues[i], err = strconv.Atoi(v)
		if err != nil {
			panic(fmt.Errorf("unable to transform value: %s", err))
		}
	}
	if caseValues[0] < minAge || caseValues[0] > maxAge || caseValues[1] < minAge || caseValues[1] > maxAge {
		panic(fmt.Errorf("case's minimum and maximum ages must be between %d and %d", minAge, maxAge))
	}
	if caseValues[2] < minInputAge || caseValues[2] > maxInputAge {
		panic(fmt.Errorf("case's input age must be between %d and %d", minInputAge, maxInputAge))
	}
	return caseValues[0], caseValues[1], caseValues[2]
}

func readNumberOfCases() int{
	n, err := strconv.Atoi(readNewLine())
	if err != nil {
		panic(fmt.Errorf("unable to get the number of cases: %s", err))
	}
	if n < minCases || n > maxCases {
		panic(fmt.Errorf("number of cases should be between %d and %d", minCases, maxCases))
	}
	return n
}

func readNewLine() string {
    text, err := reader.ReadString('\n')
    if err != nil {
        panic(fmt.Errorf("Error reading from input: %s", err))
    }
    return strings.TrimSuffix(text, "\n")
}