package main
import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    )

func main(){
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    v, err := strconv.ParseUint(scanner.Text(), 10, 32) 
    if err != nil {
        fmt.Printf("Error, not a valid number: %s", err)
        os.Exit(1)
    }
    if v > 100000 {
        fmt.Println("Error, number bigger than 100000")
        os.Exit(1)
    }
    fmt.Println(v)
}
