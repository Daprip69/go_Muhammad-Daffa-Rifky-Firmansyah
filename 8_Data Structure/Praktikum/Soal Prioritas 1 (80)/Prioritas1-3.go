// Buat program sesuai dengan deskripsi di bawah. Input merupakan variable string berisi kumpulan angka. Output merupakan list / array berisi angka yang hanya muncul 1 kali pada input.
package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	var result []int
	var condition = map[string]int{}

	for i := 0; i < len(angka); i++ {
		value := string(angka[i])
		var _, a = condition[value]
		if a {
			condition[value] = condition[value] + 1
		} else {
			condition[value] = 1
		}
		if condition[value] > 1 {
			delete(condition, value)
		}
	}
	for key := range condition {
		intKey, _ := strconv.Atoi(key)
		result = append(result, intKey)
	}
	return result
}

func main() {
	angka := "178013391010"
	fmt.Println("Angka: ", angka)
	fmt.Println("Muncul sekali: ", munculSekali(angka))
}