package models
import 
(
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ToDoList struct {
 
  ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
  Task   string             `json:"task,omitempty"`
  Status bool               `json:"status,omitempty"`
}

//Date struct
type Date struct{
	Day		int
	Month	int
	Year	int
}

//Px is the patient strut with the basic data
type Px struct{
	ID 	primitive.ObjectID
	RegNumber	int
	Name		string
	Direction	string
	Sex			bool
	IngressNumber	int
	//estadoCivil	int
	Telephone	int
	Mobile		int
	//HoraCita	string //time.Hour
	BirthDate	string //time.Date
	Age			int
	Religion	string
	Email		string
}

/*
package main

import (
	"fmt"
	"time"
)

func main() {
	format := "2006-01-02"
	input := "2018-12-22"
	s1, _ := time.Parse(format, input )
fmt.Println(s1)
	fmt.Printf("%v", s1.Format("02-01-2006"))
}
*/