package main

import "fmt"

// Function to print Fibonacci sequence up to n terms
func fibonacci(n int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a)
		a, b = b, a+b
	}
	fmt.Println() // Move to the next line after printing the sequence
}

func main() {
	var n int
	fmt.Print("Enter the number of terms in the Fibonacci sequence: ")
	fmt.Scan(&n)

	fmt.Printf("Fibonacci sequence up to %d terms: ", n)
	fibonacci(n)
}
