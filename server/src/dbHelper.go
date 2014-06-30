package main

import (
	"fmt"
	"github.com/jmcvetta/neoism"
)

var db Database

func ConnectToDb(){
	db, _ := neoism.Connect("http://localhost:7474/db/data")
}

func AddUser(username string, name string, lat float64, lon float64) {

	ConnectToDb()
	// fmt.Println(db)	
	node, _  := db.CreateNode(neoism.Props{"name":name, "username":username, "latitude":lat, "longitude": lon})
	node.AddLabel("User")
	
	//node.Delete()
	fmt.Println(node)
}

func GetUser(username string){
	ConnectToDb()
	query := neoism.CypherQuery{
		Statement: `
			MATCH (n:Person)
			WHERE n.username = {uname}
			RETURN n`,
		Parameters: neoism.Props{"uname":username}	
	}

}