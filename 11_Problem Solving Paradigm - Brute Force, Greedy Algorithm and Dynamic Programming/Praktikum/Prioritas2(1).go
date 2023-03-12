// Ada katak yang awalnya berada di atas Batu 1. Dia akan mengulangi tindakan berikut beberapa kali untuk mencapai batu N. Jika katak sedang berada di Batu i, lompat ke Batu i + 1 atau Batu i + 2. Di sini, biaya | hi - hj | terjadi, di mana j adalah batu untuk mendarat. Temukan biaya total minimum yang mungkin dikeluarkan sebelum katak mencapai Batu N.
package main
import (
	"fmt"
	"math"
)
func frog(jump []int, n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return int(math.Abs(float64(jump[0] - jump[1])))
	}
	return int(math.Min(float64(frog(jump, n-1)+int(math.Abs(float64(jump[n-1]-jump[n-2])))), float64(frog(jump, n-2)+int(math.Abs(float64(jump[n-1]-jump[n-3]))))))
}
func main() {
	fmt.Println(frog([]int{10, 30, 40, 20}, 4))         // 30
	fmt.Println(frog([]int{30, 10, 60, 10, 60, 50}, 6)) // 40
}