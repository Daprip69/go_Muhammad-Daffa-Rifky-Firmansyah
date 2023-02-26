// 1. Dalam matematika, bilangan prima adalah bilangan asli yang lebih besar dari angka 1, yang faktor pembaginya adalah 1 dan bilangan itu sendiri. 2 dan 3 adalah bilangan prima. 4 bukan bilangan prima karena 4 bisa dibagi 2. Kamu diminta untuk membuat fungsi untuk menentukan bahwa sebuah bilangan termasuk bilangan prima atau tidak.
    
// Buatlah solusi yang lebih optimal, dengan kompleksitas lebih cepat dari O(n) / O(n/2).

package main

import (
	"fmt"
	"math"
)

func primeNumber(n int) bool {
	if n < 1 {
		return false
	} 
   if n <= 3 && n >= 1{
		return true
	}
	bilangan := int(math.Sqrt(float64(n)))
	for i := 2; i <= bilangan; i += 3 {
		if n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	InputA := 1000000007
   fmt.Println("Input: 1000000007")
	if primeNumber(InputA) {
		fmt.Printf("Output for InputA: %d adalah bilangan prima.", InputA)
	} else {
		fmt.Printf("output for InputA: %d bukan bilangan prima.", InputA)
   }
      InputB := 1500450271
   fmt.Println("Input: 1500450271")
	if primeNumber(InputB) {
		fmt.Printf(" Output for Input B: %d adalah bilangan prima.", InputB)
	} else {
		fmt.Printf(" Output for Input B: %d bukan bilangan prima.", InputB)
	}
}

