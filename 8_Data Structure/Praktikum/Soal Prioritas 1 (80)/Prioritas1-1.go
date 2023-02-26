// Buatlah sebuah program menggabungkan 2 array yang diberikan, dan jangan sampai terdapat nama yang sama di data yang sudah tergabung tadi!
package main

import "fmt"

func ArrayMerge(ArrayA, ArrayB []string) []string {
	ArrayA = append(ArrayA, ArrayB...)
	result := []string{}
	condition := make(map[string]bool)

	for i := 0; i < len(ArrayA); i++ {
		condition[ArrayA[i]] = true
	}
	for key := range condition {
		result = append(result, key)
	}
	return result
}

func main() {
	ArrayA := []string{"alisa", "yoshimitsu"}
	ArrayB := []string{"devil jin", "yoshimitsu", "alisa", "law"}
	fmt.Println("ArrayA sebelum merge :", ArrayA)
	fmt.Println("ArrayB sebelum merge:", ArrayB)
	fmt.Println("Array baru setelah Merge :", ArrayMerge(ArrayA, ArrayB))
}
