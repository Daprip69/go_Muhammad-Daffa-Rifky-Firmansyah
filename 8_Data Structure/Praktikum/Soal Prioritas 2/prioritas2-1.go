// Diberi array angka yang diurutkan dan jumlah target, temukan pasangan dalam array yang jumlahnya sama dengan target yang diberikan. Tulis fungsi untuk mengembalikan indeks dari dua angka (yaitu pasangan) sehingga jika value pada index tersebut dijumlahkan sesuai target yang diberika
package main

import "fmt"

func Pairsum(arr []int, target int) []int {
	i := 0
	a := len(arr) - 1
	for i < a {
		if arr[i]+arr[a] == target {
			return []int{i, a}
		} else if arr[i]+arr[a] < target {
			i++
		} else {
			a--
		}
	}
	return []int{}
}

func main() {
	fmt.Println(Pairsum([]int{1, 2, 3, 4, 5}, 5))
	fmt.Println(Pairsum([]int{2, 5, 9, 11}, 7)) 
	fmt.Println(Pairsum([]int{1, 3, 5, 7}, 8))  
	fmt.Println(Pairsum([]int{1, 4, 6, 8}, 8)) 
	fmt.Println(Pairsum([]int{1, 5, 6, 7}, 10))    
}