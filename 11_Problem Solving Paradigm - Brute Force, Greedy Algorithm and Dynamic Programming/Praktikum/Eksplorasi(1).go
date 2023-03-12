// - Tulis program untuk mengkonversi dari angka normal ke Angka Romawi!
// Input: 4
// Output: IV
// Input: 9
// Output: IX
// Input: 23
// Output: XXIII
// Input: 2021
// Output: MMXXI
// Input: 1646
// Output: MDCXLVI
package main
import "fmt"
func konversi(num int) string {
	Romawi := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	normal := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	result := ""
	for i := 0; i < len(normal); i++ {
		for normal[i] <= num {
			result += Romawi[i]
			num -= normal[i]
		}
	}
	return result
}
func main() {
	fmt.Println(konversi(4))    // IV
	fmt.Println(konversi(9))    // IX
	fmt.Println(konversi(23))   // XXIII
	fmt.Println(konversi(2021)) // MMXXI
	fmt.Println(konversi(1646)) // MDCXLVI
}