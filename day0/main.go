package main
import (
    "fmt"
    "os"
    "bufio"
)

func main() {
    //Enter your code here. Read input from STDIN. Print output to STDOUT
    fmt.Println("Hello, World.");
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    text := scanner.Text()
    fmt.Println(text) 
}
