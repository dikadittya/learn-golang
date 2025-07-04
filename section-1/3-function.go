package main

import "fmt"

func calc(a, b int) (sum, diff int) {
	sum = a + b
	diff = a - b
	return
}

func main() {
	s, d := calc(7, 5)
	fmt.Println("Jumlah: ", s)
	fmt.Println("Selisih: ", d)
}
