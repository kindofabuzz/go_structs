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
	// fmt.Println(u.lastName)
	// fmt.Println(u.birthdate)
}

func (u *user) clearUserName() {  // must make user a pointer to change struct
	u.firstName = "deleted"
	u.lastName = "deleted"
	u.birthdate = "deleted"

}

func main() {
	userFirstName := getUserData("Enter your first name: ")
	userLastName := getUserData("Enter your last name: ")
	userBirthdate := getUserData("Enter your birthdate (MM/DD/YYYY): ")

	var appUser user

	appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthdate,
	}

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
