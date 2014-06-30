package main

import (
	// "fmt"
	"container/list"
	"github.com/jmcvetta/neoism"
	//"code.google.com/p/go-uuid/uuid"
)

const connectionString = "http://localhost:7474/db/data"

type User struct {
	name        string
	userName    string
	latitude    float64
	longitude   float64
	totalPoints int
	isActive    bool
}

type Game struct {
	name      string
	latitude  float64
	longitude float64
	isActive  bool
	Id string
}

func AddUser(username string, name string, lat float64, lon float64) {

	db, _ := neoism.Connect(connectionString)
	// fmt.Println(db)
	node, _ := db.CreateNode(neoism.Props{"name": name, "userName": username, "latitude": lat, "longitude": lon, "totalPoints": 0, "isActive": true})
	node.AddLabel("User")

	//node.Delete()
	//fmt.Println(node)
}

func AddGame(name string, lat float64, lon float64) {

	db, _ := neoism.Connect(connectionString)
	node, _ := db.CreateNode(neoism.Props{"name": name, "latitude": lat, "longitude": lon, "Id": "some guid here", "isActive": true})
	node.AddLabel("Game")

}

func GetAllGames() *list.List {
	db, _ := neoism.Connect(connectionString)
	result := []struct {
		N neoism.Node // Column "n" gets automagically unmarshalled into field N
	}{}
	query := neoism.CypherQuery{
		Statement: `
			MATCH (n:Game)
			RETURN n`,
		Result: &result,
	}
	db.Cypher(&query)
	rv := list.New()
	for _, element := range result {
		//TODO: turn this into a game
		rv.PushFront(element.N.Data)
	}
	return rv
}

func GetGame(id string) Game {
	db, _ := neoism.Connect(connectionString)
	result := []struct {
		N neoism.Node // Column "n" gets automagically unmarshalled into field N
	}{}
	query := neoism.CypherQuery{
		Statement: `
			MATCH (n:Game)
			WHERE n.Id = {id}
			RETURN n`,
		Parameters: neoism.Props{"id": id},
		Result:     &result,
	}
	db.Cypher(&query)
	rv := Game{result[0].N.Data["name"].(string), result[0].N.Data["latitude"].(float64), result[0].N.Data["longitude"].(float64), result[0].N.Data["isActive"].(bool),result[0].N.Data["Id"].(string)}
	return rv
}


func GetUser(username string) User {
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
	rv := User{result[0].N.Data["name"].(string), result[0].N.Data["userName"].(string), result[0].N.Data["latitude"].(float64), result[0].N.Data["longitude"].(float64), int(result[0].N.Data["totalPoints"].(float64)), result[0].N.Data["isActive"].(bool)}
	//fmt.Println(rv)
	return rv
}
