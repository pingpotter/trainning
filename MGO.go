package main

import (
	"fmt"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Person struct {
	Name string
	Lname string
	Phone string
	Timestamp time.Time
}

func main() {

	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("testmgo").C("people")
	err = c.Insert(&Person{"Natalie", "Portman", "+55 53 8116 9639", time.Now()},
		&Person{"Jennifer", "Lawrence ", "+55 53 8402 8510", time.Now()})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Jennifer"}).Select(bson.M{"phone": 0}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)

	// Query All
	var results []Person
	err = c.Find(bson.M{"name": "Jennifer"}).Sort("-timestamp").All(&results)

	if err != nil {
		panic(err)
	}
	fmt.Println("Results All: ", results)

	// Update
	colQuerier := bson.M{"name": "Jennifer"}
	change := bson.M{"$set": bson.M{"phone": "+86 99 8888 7777", "timestamp": time.Now()}}
	err = c.Update(colQuerier, change)
	if err != nil {
		panic(err)
	}

	// Query All
	err = c.Find(bson.M{"name": "Jennifer"}).Sort("-timestamp").All(&results)

	if err != nil {
		panic(err)
	}
	fmt.Println("Results All: ", results)

}