// Buatlah program untuk menentukan substring yang sama di antara keduanya, dengan diberikan dua string A dan B!

package main

import (
	"fmt"
	"strings"
)
func Compare(a, b string) string {
	A, B := len(a), len(b)
	if A > B {
		if strings.Contains(a, b) {
			return b
		}
	} else {
		if strings.Contains(b, a) {
			return a
		}
	}
	return ("")
}
func main() {
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
	fmt.Println(Compare("KI", "KIJANG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))
}