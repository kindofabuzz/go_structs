package main

import (
	"beast/structs/user"
	"fmt"
)

func main() {
	userFirstName := getUserData("Enter your first name: ")
	userLastName := getUserData("Enter your last name: ")
	userBirthdate := getUserData("Enter your birthdate: ")

	var appUser *user.User

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("blah@blah.com","doobie")

	admin.OutPutDetails()
	admin.ClearUserName()
	admin.OutPutDetails()

	appUser.OutPutDetails()
	appUser.ClearUserName()
	appUser.OutPutDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value

}
