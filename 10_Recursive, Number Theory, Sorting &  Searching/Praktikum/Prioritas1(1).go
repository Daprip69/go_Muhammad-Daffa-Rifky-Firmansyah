//- Dalam matematika, bilangan Fibonacci adalah barisan yang didefinisikan secara rekursif sebagai berikut:
    
//    Penjelasan: barisan ini berawal dari 0 dan 1, kemudian angka berikutnya didapat dengan cara menambahkan kedua bilangan yang berurutan sebelumnya. Dengan aturan ini, maka barisan bilangan Fibonacci yang pertama adalah:
//  0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946…..

package main
import "fmt"
func fibonacchi(num int) int{
	if num == 0 {
		return 0
	} else if num == 1 {
		return 1
	} else {
		return fibonacchi(num-1) + fibonacchi(num-2)
	}
}
func main() {
	fmt.Println(fibonacchi(0)) //output 0
	fmt.Println(fibonacchi(2)) //output 1
	fmt.Println(fibonacchi(9)) //ouput 34
	fmt.Println(fibonacchi(10)) //output 55
	fmt.Println(fibonacchi(12)) //output 144
}