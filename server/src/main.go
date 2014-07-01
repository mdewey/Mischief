package main

import (
	// "container/list"
	"fmt"

// "github.com/jmcvetta/neoism"
)

func main() {
	userName := "tester1"

	AddUser(userName, "full naming", 44.4, 55.5)
	// user := GetUser(userName)
	// fmt.Println(user)
	//DeleteAllGames()
	AddGame("Test Game 1", 43.2, 54.5)
	all := GetAllGames()
	//fmt.Println(all)
	//PrintOutList(all)

	// the first game
	firstGame := all.Front().Value.(Game)
	// fmt.Println(firstGame)
	// add user to game here
	JoinGame(userName, firstGame.code)

	// game := GetGame("some guid here")
	// fmt.Println(game)

}

// func PrintOutList(l *list.List) {
// 	if l.Len() > 0 {
// 		for e := l.Front(); e != nil; e = e.Next() {
// 			fmt.Println(e)
// 		}
// 	}

// }
