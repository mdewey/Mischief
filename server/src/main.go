package main

import (
	"fmt"

// "github.com/jmcvetta/neoism"
)

func main() {
	userName := "userNaming4"

	AddUser(userName, "full naming", 44.4, 55.5)
	user := GetUser(userName)
	fmt.Println(user)

	AddGame("New Game", 43.2, 54.5)
}
