package models

import "time"

type PersonModel struct {
	Id        int64     `bson:"_id"`
	Name      string    `bson:"name"`
	Sex       int       `bson:"sex"`
	Age       int       `bson:"age"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
