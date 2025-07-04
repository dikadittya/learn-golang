package main

import "github.com/dikadittya/learn-golang/section-2/5-oop/example-di/usecase"

func main() {
	senderUsecase := usecase.NewSmtpSenderUsecase()
	userUsecase := usecase.NewUserUsecase(senderUsecase)

	userUsecase.RegistrasiUser("dikadittya@mail.com")
}
