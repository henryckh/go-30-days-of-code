package main


import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func main() {
  	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".
  
    var i uint64 = 4
    var d float64 = 4.0
    var s string = "HackerRank "

    scanner := bufio.NewScanner(os.Stdin)
    // Declare second integer, double, and String variables.
    var incomInt uint64
    var incomDouble float64
    var incomStr string
    inputVals := []string{}
    
    // Read and save an integer, double, and String to your variables.
    for scanner.Scan() {
        input := scanner.Text()
        if len(input) == 0 {
            break
        }
        inputVals = append(inputVals, input)
    }
    
    incomInt, _ = strconv.ParseUint(inputVals[0], 10, 64)
    incomDouble, _ = strconv.ParseFloat(inputVals[1], 64)
    incomStr = inputVals[2]
    
    // Print the sum of both integer variables on a new line.
    fmt.Println(i+incomInt)
    
    // Print the sum of the double variables on a new line.
    fmt.Printf("%.1f\n", d+incomDouble)
    
    // Concatenate and print the String variables on a new line
    // The 's' variable above should be printed first.(s+incomStr)
    fmt.Println(s+incomStr)
}