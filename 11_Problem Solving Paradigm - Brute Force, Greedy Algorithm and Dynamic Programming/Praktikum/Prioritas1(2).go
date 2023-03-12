//Diberi bilangan bulat numRows, kembalikan numRows pertama dari segitiga Pascal. Dalam segitiga Pascal, setiap angka adalah jumlah dari dua angka tepat di atasnya seperti yang ditunjukkan:
package main
import "fmt"
func generatePascalTriangle(numRows int) [][]int {
    triangle := make([][]int, numRows)
    for i := 0; i < numRows; i++ {
        row := make([]int, i+1)
        for j := 0; j <= i; j++ {
            if j == 0 || j == i {
                row[j] = 1
            } else {
                row[j] = triangle[i-1][j-1] + triangle[i-1][j]
            }
        }
        triangle[i] = row
    }
    return triangle
}
func main() {
    numRows := 5
    triangle := generatePascalTriangle(numRows)
    for i := 0; i < numRows; i++ {
        fmt.Println(triangle[i])
    }
}
