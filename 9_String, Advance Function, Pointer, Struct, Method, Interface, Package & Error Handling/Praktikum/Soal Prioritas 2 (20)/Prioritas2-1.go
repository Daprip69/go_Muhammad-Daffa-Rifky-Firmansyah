// Buatlah sebuah method dengan parameter `offset` bertipe integer dan `input` bertipe string. Method ini menghasilkan sebuah string baru yang dimana setiap huruf dilakukan pergeseran berdasarkan `offset` yang merupakan jumlah pergeseran hurufnya. String `input` diasumsikan berisi huruf kecil saja dan spasi. Sebagai contoh, ketika terdapat huruf `z` yang digeser dengan `offset` sebesar 3 maka huruf hasil pergeseran adalah `c.`
//Daftar karakter ASCII dapat dilihat di link [berikut](https://id.wikipedia.org/wiki/ASCII).
// Berdasarkan referensi ASCII, huruf `a` memiliki kode 97, huruf `b` memiliki kode 98, `z` memiliki kode 122. Anda bisa menggunakan operator modulo jika diperlukan.

//Test Case 1

//Input: `offset = 3` , `input = abc`

//Output: `def`

//Test Case 2

//Input: `offset = 1`, `input = abcdefghijklmnopqrstuvwxyz`

//Output: `bcdefghijklmnopqrstuvwxyza`

//Test Case 3

//Input: `offset = 1000`, `input = abcdefghijklmnopqrstuvwxyz`

//Output: `mnopqrstuvwxyzabcdefghijkl`

//Test Case lain dapat dicoba di [https://cryptii.com/](https://cryptii.com/)
package main

import (
	"fmt"
)

func caesar(offset int, input string) string {
	result := []rune{}
	for _, char := range input {
		char = 'a' + (char-'a'+rune(offset))%26
		result = append(result, char)
	}
	return string(result)
}

func main() {
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(2, "alta"))                      
	fmt.Println(caesar(10, "alterraacademy"))
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))
}