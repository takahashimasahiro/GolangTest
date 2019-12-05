package main


import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func PasswordCreate() {
	password := "password"
	fmt.Println(password)
	vytepass := []byte(password)
	fmt.Println(vytepass)
	// hash化(byte slice)
	converted, err:= bcrypt.GenerateFromPassword(vytepass, 10)
	if err != nil {
		fmt.Println(err)
		return 
	}
	fmt.Println(converted)
	// cast(byte slice to string)
	fmt.Println(string(converted))
}