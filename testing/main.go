package main

import (
	"fmt"
	"strings"
)

func main() {
	email := "some@hotmail.com"

	emailStripedHash := strings.Split(email, "#")

	email = emailStripedHash[len(emailStripedHash)-1]

	fmt.Println(email)

}
