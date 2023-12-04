package main

import "fmt"

func main() {
fmt.Println(cal(2,3))
fmt.Println(createName("desmond", "Dikas"))
fmt.Println(factoria(10))
} 

// A function that sums two different interger Numbers
func cal(x int, y int)int{
	return x + y
}

// A function that concatenates two strings

func createName(firstName string, secondName string) string {
	// fullName := fmt.Sprintf("%v %v", firstName, secondName)
	fullName := firstName+ " "+ secondName 
	return fullName
}

func factoria(n int) int {
	f := n
	for i := 1; i < n; i++ {
		f *= (n - i)
	}
	return f
}