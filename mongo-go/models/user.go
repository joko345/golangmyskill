package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id     primitive.ObjectID `bson:"_id,omitempty"` //memakai undescore sebagai penanda primaryKey
	Name   string             `json:"name" bson:"name"`
	Gender string             `json:"gender" bson:gender"`
	Age    int                `json:"age" bson:age"`
}

func main() {

}
