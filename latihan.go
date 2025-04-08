package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func printPascal(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%d ", combination(i, j))
		}
		fmt.Println()
	}
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&n)
	printPascal(n)
}
