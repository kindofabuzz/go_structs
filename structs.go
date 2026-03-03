package main

import (
	"fmt"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
}

func (u user) outPutDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)

}

func (u *user) clearUserName() { // must make user a pointer to change struct
	u.firstName = "deleted"
	u.lastName = "deleted"
	u.birthdate = "deleted"

}

func newUser(firstName, lastName, birthdate string) *user {
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
	}
}

func main() {
	userFirstName := getUserData("Enter your first name: ")
	userLastName := getUserData("Enter your last name: ")
	userBirthdate := getUserData("Enter your birthdate: ")

	var appUser *user

	appUser = newUser(userFirstName, userLastName, userBirthdate)

	appUser.outPutDetails()
	appUser.clearUserName()
	appUser.outPutDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value

}
