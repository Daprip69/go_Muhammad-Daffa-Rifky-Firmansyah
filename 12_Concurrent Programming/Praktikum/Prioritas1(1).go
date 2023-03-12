//Buatlah sebuah fungsi yang mencetak angka kelipatan x, dimana x merupakan parameter bilangan bulat positif. lalu jalankan fungsi tersebut dengan menerapkan goroutine, dengan interval waktu 3 detik!
package main
import (
    "fmt"
    "time"
)
func multiple(x int) {
    for i := 1; i <= 10; i++ {
        fmt.Printf("%d x %d = %d\n", i, x, i*x)
        time.Sleep(3 * time.Second)
    }
}
func main() {
    go multiple(3)
    go multiple(2)
    go multiple(69)
    time.Sleep(35 * time.Second)
}
