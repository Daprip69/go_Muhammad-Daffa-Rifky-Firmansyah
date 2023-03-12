//Dalam matematika, bilangan prima adalah bilangan asli yang lebih besar dari angka 1, yang faktor pembaginya adalah 1 dan bilangan itu sendiri. Angka 2 dan 3 adalah bilangan prima. Angka 4 bukan bilangan prima karena 4 bisa dibagi 2. Sepuluh deret bilangan prima yang pertama adalah 2, 3, 5, 7, 11, 13, 17, 19, 23 dan 29. Buatlah sebuah fungsi bernama getPrime yang menampilkan bilangan prima sesuai dengan deret urutannya.
package main
import "fmt"
func isPrime(num int) bool {
    if num < 2 {
        return false
    }
    for i := 2; i*i <= num; i++ {
        if num%i == 0 {
            return false
        }
    }
    return true
}
func primeX(n int) int {
    count := 0
    num := 2
    for count < n {
        if isPrime(num) {
            count++
        }
        num++
    }
    return num-1
}
func main() {
    fmt.Println(primeX(1))  // 2
    fmt.Println(primeX(5))  // 11
    fmt.Println(primeX(8))  // 19
    fmt.Println(primeX(9))  // 23
    fmt.Println(primeX(10)) // 29
}
