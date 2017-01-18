package database

import (
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/satori/go.uuid"
	"github.com/hellaesir/entities"
)

func AddTask(obj entities.Task) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
			panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("todo").C("task")
	//err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
	//           &Person{"Cla", "+55 53 8402 8510"})
	
	obj.Id = uuid.NewV4().String()
	
	err = c.Insert(obj)
			   
	if err != nil {
			log.Fatal(err)
	}
}

func GetTask(id string) entities.Task{
	session, err := mgo.Dial("localhost:27017")
	c := session.DB("todo").C("task")
	
	result := entities.Task{}
	
	err = c.Find(bson.M{"Id": id}).One(&result)
	
	if err != nil {
		log.Fatal(err)
	}
	
	return result;
}

func GetAllTasks() []entities.Task{
	session, err := mgo.Dial("localhost:27017")

	c := session.DB("todo").C("task")
	
	var results []entities.Task
	
	err = c.Find(nil).All(&results)
	
	if err != nil {
		log.Fatal(err)
	}
	
	return results;
}