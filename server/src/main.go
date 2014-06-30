package main

import (
	"container/list"
	"fmt"

// "github.com/jmcvetta/neoism"
)

func main() {
	//	userName := "userNaming4"

	//AddUser(userName, "full naming", 44.4, 55.5)
	// user := GetUser(userName)
	// fmt.Println(user)
	//AddGame("New Game 2", 43.2, 54.5)
	all := GetAllGames()
	fmt.Println(all)
	PrintOutList(all)

	game := GetGame("some guid here")
	fmt.Println(game)

}

func PrintOutList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e)
	}
}
