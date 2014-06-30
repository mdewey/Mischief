package main

import (
	"fmt"
	"github.com/jmcvetta/neoism"
)

func AddUser(username string, name string, lat float64, lon float64) {

	db, _ := neoism.Connect("http://localhost:7474/db/data")
	// fmt.Println(db)	
	node, _  := db.CreateNode(neoism.Props{"name":name, "username":username, "latitude":lat, "longitude": lon})
	node.AddLabel("User")
	
	//node.Delete()
	fmt.Println(node)
}

func GetUser(username string){
	
}