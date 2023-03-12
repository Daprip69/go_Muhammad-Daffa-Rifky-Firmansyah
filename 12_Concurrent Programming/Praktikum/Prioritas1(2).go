//Buatlah program yang mencetak bilangan kelipatan 3 dengan menerapkan goroutine dan channel!
package main
import (
    "fmt"
    "time"
)
func generateMultiples(ch chan<- int) {
    for i := 1; i <= 1000; i++ {
        ch <- i * 3
    }
    close(ch)
}
func printMultiples(ch <-chan int) {
    for multiple := range ch {
        fmt.Println(multiple)
    }
}
func main() {
    ch := make(chan int)
    go generateMultiples(ch)
    go printMultiples(ch)
    time.Sleep(time.Second) // Block the main goroutine
}
