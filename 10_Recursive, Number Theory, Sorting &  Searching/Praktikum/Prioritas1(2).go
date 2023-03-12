// Buatlah program di Golang yang dapat mengurutkan barang berdasarkan jumlah kemunculannya. Jika ada barang yang duplicate kamu hanya perlu memunculkan sekali, namun kamu perlu menampilkan total kemunculan barang tersebut!
package main
import (
	"fmt"
	"sort"
)
type Pair struct {
	name  string
	count int
}
func MostAppearItem(items []string) []Pair {
	Counts := make(map[string]int)
	for _, item := range items {
		if _, ok := Counts[item]; ok {
			Counts[item]++
		} else {
			Counts[item] = 1
		}
	}
	var result []Pair
	for key, val := range Counts {
		result = append(result, Pair{name: key, count: val})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].count > result[j].count
	})
	return result
}
func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
}
