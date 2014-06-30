package main

import (
	//"fmt"
	"github.com/jmcvetta/neoism"
)

const connectionString = "http://localhost:7474/db/data"

type User struct{
	name string
	userName string
	latitude float64
	longitude float64
	totalPoints int
	isActive bool
}

func AddUser(username string, name string, lat float64, lon float64) {

	db, _ := neoism.Connect(connectionString)
	// fmt.Println(db)
	node, _ := db.CreateNode(neoism.Props{"name": name, "userName": username, "latitude": lat, "longitude": lon, "totalPoints": 0, "isActive": true})
	node.AddLabel("User")

	//node.Delete()
	//fmt.Println(node)
}

func GetUser(username string) User{
	db, _ := neoism.Connect(connectionString)
	result := []struct {
		N neoism.Node // Column "n" gets automagically unmarshalled into field N
	}{}
	query := neoism.CypherQuery{
		Statement: `
			MATCH (n:User)
			WHERE n.userName = {userName}
			RETURN n`,
		Parameters: neoism.Props{"userName": username},
		Result:     &result,
	}
	db.Cypher(&query)
	rv := User{result[0].N.Data["name"].(string),result[0].N.Data["userName"].(string),result[0].N.Data["latitude"].(float64),result[0].N.Data["longitude"].(float64),int(result[0].N.Data["totalPoints"].(float64)),result[0].N.Data["isActive"].(bool)}
	//fmt.Println(rv)	
	return rv
}
