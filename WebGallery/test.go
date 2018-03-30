package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func main() {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		log.Fatal(err)
		return
	}

	defer session.Close()
	collection := session.DB("test").C("go")
	err = collection.Insert(bson.M{"name": "Pradit"})
	if err != nil {
		log.Fatal(err)
		return
	}
}
