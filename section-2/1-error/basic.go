package main

import (
	"errors"
	"fmt"
)

func validasiUmur(umur int) error {
	if umur < 0 {
		return errors.New("UMUR TIDAK BOLEH NEGATIF") // string error disarankan besar semua
	}

	if umur > 10 {
		return fmt.Errorf("UMUR TIDAK BOLEH MELEBIHI 10, INPUTAN ANDA : %d", umur)
	}

	return nil
}

func main() {
	err := validasiUmur(5)
	if err != nil {
		fmt.Println("Error: ", err)
		// input log error
		return
	}

	fmt.Println("Berhasil")
	return
}
