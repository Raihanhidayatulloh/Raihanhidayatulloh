package main

import "fmt"

func main() {
	var x float32

	fmt.Println("Tolong masukan angka untuk var x apa saja:")
	fmt.Scan(&x)

	var fx = ((x * x * x) + (3 * x)) / ((x * x * x * x) - (3 * x * x) + 4)
	fmt.Println(fx)
}

