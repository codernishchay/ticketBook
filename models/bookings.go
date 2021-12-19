package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Booking struct {
	ShowID      primitive.ObjectID `bson:"showid" json:"showid"`
	Ticketprice int                `bson:"ticketprice" json:"ticketprice"`
	Showtime    string             `bson:"showtime" json:"showtime"`
	UserID      primitive.ObjectID `bson:"userid"`
	Username    string             `bson:"username" json:"username"`
	Seat        Seat               `bson:"seat" json:"seat"`
}
