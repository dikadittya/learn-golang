package main

import "fmt"

func main() {
	age := 18
	if age < 17 {
		fmt.Println("Belum cukup umur")
	} else if age == 17 {
		fmt.Println("Pas 17 tahun")
	} else {
		fmt.Println("Sudah dewasa")
	}

	fmt.Println("--------------")

	day := "Senin"

	switch day {
	case "Senin":
		fmt.Println("Hari kerja dimulai")
	case "Sabtu", "Minggu":
		fmt.Println("Weekend!")
	default:
		fmt.Println("Hari biasa")
	}

	fmt.Println("--------------")

	switch {
	case age < 17:
		fmt.Println("Remaja")
	case age < 30:
		fmt.Println("Dewasa Muda")
	default:
		fmt.Println("Dewasa")
	}
}
