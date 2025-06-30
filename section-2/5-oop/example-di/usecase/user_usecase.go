package usecase

import (
	"fmt"
)

// class User Service
type UserUsecase struct {
	sender ISmtpSenderUsecase // dependency
}

// Constructor
func NewUserUsecase(sender ISmtpSenderUsecase) *UserUsecase {
	return &UserUsecase{sender: sender}
}

// Interface abstract untuk UserUsecase (polymorphism)
type IUserUsecase interface {
	RegistrasiUser(email string)
}

// Implementasi method dari interface IUserUsecase
func (u *UserUsecase) RegistrasiUser(email string) {
	fmt.Println("Registrasu user berhasil:", email)
	u.sender.KirimEmail(email, "Selamat datang di platform kami!")
}
