// Buat sebuah struct dengan nama Student yang mempunyai properti name dan score dalam bentuk slice kemudian simpan data siswa sebanyak 5 siswa. Setelah 5 siswa dimasukkan maka program menunjukkan skor rata-rata, siswa yang memiliki skor minimum dan maksimal? (implementasikan method)

package main

import (
	"fmt"
	"math"
)
type Student struct {
	Name  string
	Score int
}
type Students []Student
func (s Students) AverageScore() float64 {
	var sum int
	for _, student := range s {
		sum += student.Score
	}
	return float64(sum) / float64(len(s))
}
func (s Students) minimum() (string, int) {
	var minScore int = math.MaxInt32
	var minName string
	for _, student := range s {
		if student.Score < minScore {
			minScore = student.Score
			minName = student.Name
		}
	}
	return minName, minScore
}
func (s Students) maximum() (string, int) {
	var maxScore int = math.MinInt32
	var maxName string
	for _, student := range s {
		if student.Score > maxScore {
			maxScore = student.Score
			maxName = student.Name
		}
	}
	return maxName, maxScore
}
func main() {
	var students Students
	for i := 0; i < 5; i++ {
		var name string
		var score int
		fmt.Printf("Input %d student name: ", i+1)
		fmt.Scan(&name)
		fmt.Printf("Input %d student score: ", i+1)
		fmt.Scan(&score)
		student := Student{
			Name:  name,
			Score: score,
		}
		students = append(students, student)
	}
	averageScore := students.AverageScore()
	minName, minScore := students.minimum()
	maxName, maxScore := students.maximum()
	fmt.Printf("Average score: %.2f\n", averageScore)
	fmt.Printf("Min score of student: %s %d\n", minName, minScore)
	fmt.Printf("Max score of student: %s %d\n", maxName, maxScore)
}
