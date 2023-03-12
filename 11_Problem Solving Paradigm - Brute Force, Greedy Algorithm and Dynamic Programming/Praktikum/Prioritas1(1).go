//- Diberi bilangan bulat n, kembalikan array ans dengan panjang n + 1 sehingga untuk setiap i (0 <= i <= n), ans[i] adalah bilangan 1 dalam representasi biner dari i
//Input: n = 2  
//Output: [0,1,10]
//Input: n = 3
//Output: [0,1,10, 11]
package main
import "fmt"
func ansbiner(n int) {
	for i := 0; i <= n; i++ {
		fmt.Printf("%b \n", i)
	}
}
func main() {
	ansbiner(2)
	ansbiner(3)
}