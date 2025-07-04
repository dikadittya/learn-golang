package main

import "fmt"

type Mahasiswa struct {
	Nama    string
	Umur    int
	Jurusan string
}

func main() {
	mahasiswaA := Mahasiswa{
		Nama:    "Samson",
		Umur:    20,
		Jurusan: "TI",
	}

	var mahasiswaB Mahasiswa
	mahasiswaB.Nama = "Tarzan"
	mahasiswaB.Umur = 24
	mahasiswaB.Jurusan = "Guru"

	mahasiswaC := Mahasiswa{"Andi", 22, "Dokter"}

	fmt.Println("Mahasiswa A:", mahasiswaA)
	fmt.Println("Mahasiswa B:", mahasiswaB)
	fmt.Println("Mahasiswa C:", mahasiswaC)

	fmt.Println("Nama Mahasiswa A:", mahasiswaA.Nama)

	// update data
	mahasiswaA.Jurusan = "Manajemen"
	fmt.Println("Jurusan terbaru mahasiswa A", mahasiswaA.Jurusan)

	// Gunakan struct sbg parameter fungsi
	printMahasiswa(mahasiswaB)
}

func printMahasiswa(s Mahasiswa) {
	fmt.Printf("Nama : %s\nUsia: %d\nJurusan: %s\n", s.Nama, s.Umur, s.Jurusan)
}
