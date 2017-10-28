package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Person struct {
	Name      string    `bson:"name"`
	Phone     string    `bson:"phone"`
	CreatedAt time.Time `bson:"createdAt"`
}

const (
	IsDrop = true
)

func main() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	// Drop Database
	if IsDrop {
		err = session.DB("test").DropDatabase()
		if err != nil {
			panic(err)
		}
	}

	// Collection People
	c := session.DB("test").C("people")

	// Index
	index := mgo.Index{
		Key: []string{"createdAt"},
		//Unique:      true,
		//DropDups:    true,
		Background:  true,
		//Sparse:      true,
		ExpireAfter: time.Second * 24,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	// Insert Datas
	err = c.Insert(&Person{Name: "Ale", Phone: "+55 53 1234 4321", CreatedAt: time.Now()},
		&Person{Name: "Cla", Phone: "+66 33 1234 5678", CreatedAt: time.Now()})

	if err != nil {
		panic(err)
	}

	// Query One
	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).Select(bson.M{"phone": 0}).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println("Phone", result)

	// Query All
	var results []Person
	err = c.Find(bson.M{"name": "Ale"}).Sort("-createdAt").All(&results)

	if err != nil {
		panic(err)
	}
	fmt.Println("Results All: ", results)

	// Update
	query := bson.M{"name": "Ale"}
	change := bson.M{"$set": bson.M{"phone": "+86 99 8888 7777"}}
	err = c.Update(query, change)
	if err != nil {
		panic(err)
	}

	// Query All
	err = c.Find(bson.M{"name": "Ale"}).Sort("-createdAt").All(&results)

	if err != nil {
		panic(err)
	}
	fmt.Println("Results All: ", results)

	// Query All
	err = c.Find(nil).Sort("-createdAt").All(&results)

	if err != nil {
		panic(err)
	}
	fmt.Println("Results All: ", results)
}
