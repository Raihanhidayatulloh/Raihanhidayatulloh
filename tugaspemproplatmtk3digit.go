package main
import "fmt"
func main() {
	var x int

	fmt.Print("Masukkan bilangan bulat positif (<= 999): ")
	fmt.Scan(&x)

	if x <= 0 || x > 999 {
		fmt.Println("Masukkan bilangan bulat positif kurang dari 999 yang valid.")
		return
	} else {
		d1 := x / 100
		d2 := (x / 10) % 10
		d3 := x % 10

		fmt.Printf("Digit pertama: %d\n", d1)
		fmt.Printf("Digit kedua: %d\n", d2)
		fmt.Printf("Digit ketiga: %d\n", d3)
	}
}
