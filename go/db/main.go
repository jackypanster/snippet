package main

import (
	//"fmt"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"time"
	"github.com/jackypanster/util"
)

type Person struct {
	Name      string    `bson:"name"`
	Phone     string    `bson:"phone"`
	CreatedAt time.Time `bson:"createdAt"`
}

const (
	IsDrop = true
)

func init() {
	util.InitQueue(512, 65536)
}

func main() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	/*
	initDb()

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
	*/

	// Drop Database
	//initDb(session)
	// Collection People
	//c := session.DB("test").C("people")
	//Insert(c)

	// Drop Database
	initDb(session)
	// Collection People
	c := session.DB("test").C("people")
	InsertBatch(c)

	// Drop Database
	//initDb(session)
	// Collection People
	//c = session.DB("test").C("people")
	//InsertConcurrency(c)

	initDb(session)
	// Collection People
	c = session.DB("test").C("people")
	InsertConcurrencyBatch(c)
}

func initDb(session *mgo.Session) {
	if IsDrop {
		err := session.DB("test").DropDatabase()
		if err != nil {
			panic(err)
		}
	}

	c := session.DB("test").C("people")

	// Index
	index := mgo.Index{
		Key: []string{"createdAt"},
		//Unique:      true,
		//DropDups:    true,
		Background: true,
		//Sparse:      true,
		ExpireAfter: time.Hour * 1,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	/*err = c.Create(&mgo.CollectionInfo{
		Capped:   true,
		MaxBytes: 1073741824,
	})
	if err != nil {
		panic(err)
	}*/
}
