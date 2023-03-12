//Angka Fibonacci adalah serangkaian angka di mana setiap angka adalah jumlah dari keduanya nomor sebelumnya. Beberapa angka Fibonacci pertama adalah: 0, 1, 1, 2, 3, 5, 8, ….
//Buatlah fungsi untuk menghitung angka Fibonacci ke-n (top-down)!
package main
import "fmt"
func fibonacci(limit int) []int {
    a, b := 0, 1
    fibo := []int{}
    for a <= limit {
        fibo = append(fibo, a)
        a, b = b, a+b
    }
    return fibo
}
func main() {
    fmt.Println(fibonacci(0))  // 0
	fmt.Println(fibonacci(1))  // 1
	fmt.Println(fibonacci(2))  // 1
	fmt.Println(fibonacci(3))  // 2
	fmt.Println(fibonacci(5))  // 5
	fmt.Println(fibonacci(6))  // 8
	fmt.Println(fibonacci(7))  // 13
	fmt.Println(fibonacci(9))  // 34
	fmt.Println(fibonacci(10)) // 55
}
