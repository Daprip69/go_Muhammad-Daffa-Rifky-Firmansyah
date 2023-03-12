//1. Pada problem ini kamu harus menemukan total maksimum jumlah bilangan dari deret sebuah integer secara berurutan.
//Sample Test Case
//Input: [-2, 1, -3, 4, -1, 2, 1, -5, 4]
//Output: 6
//**Penjelasan: 6** adalah hasil penambahan dari deret 4, -1, 2, 1
//Sample Test Case
//Input: [-2, -5, 6, -2, -3, 1, 5, -6]
//Output: 7
//************************Penjelasan:************************ 7 adalah hasil penambahan dari deret 6, -2, -3, 1, 5
package main
import "fmt"
func MaxSequence(arr []int) int {
	SoFarmax, EndingHere := 0, 0
	for _, num := range arr {
		EndingHere += num
		if EndingHere < 0 {
			EndingHere = 0
		}
		if EndingHere > SoFarmax {
			SoFarmax = EndingHere
		}
	}
	return SoFarmax
}
func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))  // 6
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))    // 7
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))    // 7
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))    // 8
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))     // 12
}
