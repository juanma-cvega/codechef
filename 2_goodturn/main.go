package main 

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

var reader *bufio.Reader

func main(){
    reader = bufio.NewReader(os.Stdin)
	tries := readTries()
	for i := 0;i < tries;i++ {
	    diceValues := readDiceValues()
	    writeResultFrom(isGood(diceValues))
	}
}

func readTries() int {
    l := readNewLine()
    v, err := strconv.Atoi(l)
    if err != nil {
        panic(fmt.Errorf("Error while transforming input into number: %s", err))
    }
    if v < 1 && v > 100 {
        panic(fmt.Errorf("Number of tries is not within range (1 - 100)"))
    }
    return v
}

func readDiceValues() []int{
    l := readNewLine()
    vv := strings.Split(l, " ")
    if len(vv) != 2 {
        panic(fmt.Errorf("Number of dice thrown invalid, it should be 2. Current: %d", len(vv)))
    }
    dice := make([]int, len(vv))
    for i, v := range vv {
        dice[i], _ = strconv.Atoi(v)
    }
    if (dice[0] < 1 && dice[0] > 6) || (dice[1] < 1 && dice[1] > 6) {
        panic(fmt.Errorf("Dice values should be between 1 and 6."))
    }
    return dice
}

func isGood(dice []int) string {
    if dice[0] + dice[1] > 6 {
        return "YES"
    }
    return "NO"
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