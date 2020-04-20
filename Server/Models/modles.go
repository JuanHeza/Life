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
	//HoraCita	time.Duration
	BirthDate	Date
	Age			int
	Religion	string
	Email		string
}

/*

*/