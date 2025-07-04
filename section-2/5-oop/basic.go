package main

import "fmt"

type Account struct {
	nama     string
	Username string
}

func (u *Account) SetNama(n string) {
	u.nama = n
}

func (u Account) GetNama() string {
	return u.nama
}

func main() {
	u := Account{Username: "haryono"}
	u.SetNama("Haryono Putro")

	fmt.Println("Username:", u.Username)
	fmt.Println("Nama Lengkap:", u.GetNama())
	fmt.Println(u.nama)
}
