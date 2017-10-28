package main

import (
	"log"
	"time"
	"gopkg.in/mgo.v2"
	"github.com/jackypanster/util"
)

const (
	BatchSize = 100
	Amount    = 10000 * 1000
)

func batch() []interface{} {
	var persons []interface{}
	for i := 0; i < BatchSize; i ++ {
		persons = append(persons, &Person{Name: "Ale", Phone: "+55 53 1234 4321", CreatedAt: time.Now()})
	}
	return persons
}

func FindAll(c *mgo.Collection) {
	start := time.Now()
	it := c.Find(nil).Iter()
	var person Person
	idx := 0
	for it.Next(&person) {
		idx++
		log.Printf("%d %v\n", idx, person)
	}
	log.Printf("DURATION %s\n", time.Now().Sub(start))
}

func Insert(c *mgo.Collection) {
	//log.Printf("BEGIN %s\n", time.Now())
	start := time.Now()
	for i := 0; i < Amount; i++ {
		err := c.Insert(&Person{Name: "Ale", Phone: "+55 53 1234 4321", CreatedAt: time.Now()})
		if err != nil {
			panic(err)
		}
	}
	log.Printf("DURATION %s\n", time.Now().Sub(start))
}

func InsertBatch(c *mgo.Collection) {
	//log.Printf("BEGIN %s\n", time.Now())
	start := time.Now()
	for i := 0; i < (Amount / BatchSize); i++ {
		err := c.Insert(batch()...)
		if err != nil {
			panic(err)
		}
	}
	log.Printf("DURATION %s\n", time.Now().Sub(start))
}

func InsertConcurrency(c *mgo.Collection) {
	done := make(chan bool)
	//log.Printf("BEGIN %s\n", time.Now())
	start := time.Now()
	for i := 0; i < Amount; i++ {
		util.JobQueue <- util.Job{
			Do: func() error {
				err := c.Insert(&Person{Name: "Ale", Phone: "+55 53 1234 4321", CreatedAt: time.Now()})
				if err != nil {
					done <- false
					return err
				}
				done <- true
				return nil
			},
		}
	}
	//log.Printf("RUNNING %s\n", time.Now())
	log.Printf("RUNNING %s\n", time.Now().Sub(start))
	for i := 1; i <= Amount; i++ {
		select {
		case <-done:
			if i == 200000 {
				log.Println("20%")
			}
			if i == 400000 {
				log.Println("40%")
			}
			if i == 600000 {
				log.Println("60%")
			}
			if i == 800000 {
				log.Println("80%")
			}
			if i == 1000000 {
				log.Println("100%")
			}
		}
	}
	log.Printf("DURATION %s\n", time.Now().Sub(start))
}

func InsertConcurrencyBatch(c *mgo.Collection) {
	done := make(chan bool, Amount/BatchSize)
	//log.Printf("BEGIN %s\n", time.Now())
	start := time.Now()
	for i := 0; i < Amount/BatchSize; i++ {
		util.JobQueue <- util.Job{
			Do: func() error {
				err := c.Insert(batch()...)
				if err != nil {
					done <- false
					return err
				}
				done <- true
				return nil
			},
		}
	}
	//log.Printf("RUNNING %s\n", time.Now())
	log.Printf("RUNNING %s\n", time.Now().Sub(start))
	for i := 1; i <= Amount/BatchSize; i++ {
		select {
		case <-done:
			if i == 20000 {
				log.Println("20%")
			}
			if i == 40000 {
				log.Println("40%")
			}
			if i == 60000 {
				log.Println("60%")
			}
			if i == 80000 {
				log.Println("80%")
			}
			if i == 100000 {
				log.Println("100%")
			}
		}
	}
	log.Printf("DURATION %s\n", time.Now().Sub(start))
}
