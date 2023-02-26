//buatlah sebuah program yang dapat menghitung berapa banyak sebuah string yang sama didalam sebuah slice!
package main

import "fmt"

func Mapping(slice []string) map[string]int {
	var result = map[string]int{}

	for i := 0; i < len(slice); i++ {
		var _, exist = result[slice[i]]
		if exist {
			result[slice[i]] = result[slice[i]] + 1
		} else {
			result[slice[i]] = 1
		}
	}
	return result
}

func main() {
	slice := []string{"king", "devil jin", "akuma", "eddie", "steve", "geese", "king", "devil jin", "akuma", "eddie", "steve", "geese"}
	fmt.Println("Before :", slice)
	fmt.Println(Mapping(slice))
}