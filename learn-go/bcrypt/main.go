package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "DrowRanger#651"

	fmt.Println([]byte(password))

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 4)

	if err != nil {
		fmt.Println("Problem with generating password encryption")
	}

	fmt.Println(string(encryptedPassword))

	loginPassword := "DrowRanger#6a1"

	err = bcrypt.CompareHashAndPassword(encryptedPassword, []byte(loginPassword))

	if err != nil {
		fmt.Println("Loging failed, password incorrect")
		return
	}

	fmt.Println("You're logged in")
}
