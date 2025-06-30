package main

import "fmt"

type Mahasiswa struct {
	Nama    string
	Umur    string
	Jurusan string
}

func main() {
	mahasiswaA := Mahasiswa{
		Nama:    "Samson",
		Umur:    "20",
		Jurusan: "TI",
	}

	fmt.Println(mahasiswaA)

	// update data
	mahasiswaA.Jurusan = "Manajemen"
	fmt.Println(mahasiswaA)
}
