package main

import (
	"fmt"
	"github.com/jmcvetta/neoism"
)

const connectionString = "http://localhost:7474/db/data"

func AddUser(username string, name string, lat float64, lon float64) {

	db, _ := neoism.Connect(connectionString)
	// fmt.Println(db)
	node, _ := db.CreateNode(neoism.Props{"name": name, "userName": username, "latitude": lat, "longitude": lon, "totalPoints": 0, "isActive": true})
	node.AddLabel("User")

	//node.Delete()
	//fmt.Println(node)
}

func GetUser(username string) {
	db, _ := neoism.Connect(connectionString)
	rv := []struct {
		N neoism.Node // Column "n" gets automagically unmarshalled into field N
	}{}
	query := neoism.CypherQuery{
		Statement: `
			MATCH (n:User)
			WHERE n.userName = {userName}
			RETURN n`,
		Parameters: neoism.Props{"userName": username},
		Result:     &rv,
	}
	db.Cypher(&query)
	fmt.Println(rv[0].N.Data["userName"])
	//return rv
}
