package main

import "fmt"

func main() {
	tanpaPointer := 3000

	var denganPointer *int
	denganPointer = &tanpaPointer

	fmt.Println("Nilai tanpa Pointer", tanpaPointer)          // Nilai tanpa Pointer 3000
	fmt.Println("Nilai dengan Pointer", denganPointer)        // Nilai dengan Pointer 0xc00000a0d8
	fmt.Println("Akses Nilai dengan Pointer", *denganPointer) // Akses Nilai dengan Pointer 3000

	luas, err := hitungLuasPersegiPanjang(10, 5)
	if err != nil {
		fmt.Println("Error : ", *luas)
	}
	fmt.Println("Luas : ", *luas)
}

func hitungLuasPersegiPanjang(panjang, lebar int) (*int, error) {
	if panjang == 0 || lebar == 0 {
		return nil, fmt.Errorf("PANJANG DAN LEBAR TIDAK BOLEH 0!")
	}

	luas := panjang * lebar
	return &luas, nil
}
