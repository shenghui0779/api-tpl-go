package models

import "time"

type Student struct {
	ID        int       `bson:"_id"`
	Name      string    `bson:"name"`
	Sex       string    `bson:"sex"`
	Age       int       `bson:"age"`
	School    string    `bson:"school"`
	Grade     string    `bson:"grade"`
	Class     string    `bson:"class"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
