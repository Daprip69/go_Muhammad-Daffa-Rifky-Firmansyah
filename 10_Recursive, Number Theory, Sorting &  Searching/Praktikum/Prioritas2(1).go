//Buatlah program playingDomino yang menerima 2 parameter array; parameter pertama merupakan kartu domino yang ada di tangan, • Parameter kedua merupakan kartu yang sedang ada di deck. Jika ada kartu yang disarankan maka output: [x,y], jika tidak ada kartu yang sesuai maka keluarkan: 'tutup kartu'.
package main
import "fmt"
func playingDomino(card [][]int, deck []int) interface{} {
	for _, val := range card {
		if val[0] == deck[0] || val[0] == deck[1] || val[1] == deck[0] || val[1] == deck[1] {
			return val
		}
	}
	return "Tutup kartu"
}
func main() {
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 4}, {2, 1}, {3, 3}}, []int{4, 3}))
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 3}, {3, 4}, {2, 1}}, []int{3, 6}))
	fmt.Println(playingDomino([][]int{{6, 6}, {2, 4}, {3, 6}}, []int{5, 11}))
}