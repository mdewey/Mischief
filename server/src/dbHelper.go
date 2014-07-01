package main

import (
	"container/list"
	"fmt"
	"github.com/jmcvetta/neoism"
	"github.com/nu7hatch/gouuid"
	"math/rand"
	"strings"
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
	Id        string
	code      string
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
	newId, _ := uuid.NewV4()
	fmt.Println(newId)
	//genCode()
	node, _ := db.CreateNode(neoism.Props{"name": name, "latitude": lat, "longitude": lon, "Id": newId.String(), "isActive": true, "code": "ABCD"})
	node.AddLabel("Game")

}

func genCode() {
	chars := strings.Split("A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P,Q,R,S,T,U,V,W,X,Y,Z,1,2,3,4,5,6,7,8,9,0", ",")
	code := list.New()
	for i := 0; i < 7; i++ {
		pos := rand.Intn(len(chars))
		code.PushFront(chars[pos])
	}
	fmt.Println(code)
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
		rv.PushFront(Game{element.N.Data["name"].(string), element.N.Data["latitude"].(float64), element.N.Data["longitude"].(float64), element.N.Data["isActive"].(bool), element.N.Data["Id"].(string), element.N.Data["code"].(string)})
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
	var rv Game
	if len(result) > 0 {
		rv = Game{result[0].N.Data["name"].(string), result[0].N.Data["latitude"].(float64), result[0].N.Data["longitude"].(float64), result[0].N.Data["isActive"].(bool), result[0].N.Data["Id"].(string), result[0].N.Data["code"].(string)}
	}	
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

func DeleteAllGames() {
	db, _ := neoism.Connect(connectionString)
	query := neoism.CypherQuery{
		Statement: `
			MATCH (n:Game)
			DELETE n`,
	}
	db.Cypher(&query)

}
