package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Show struct {
	ID          primitive.ObjectID `bson:"_id" omitempty json:"_id"`
	Time        string             `bson:"title" json:"title"`
	Title       string             `bson:"title" json:"title"`
	Pricing     int                `bson:"pricing" json:"pricing"`
	Image       []string           `bson:"images" json:"image"`
	Description string             `bson:"description" json:"description"`
	Host        Host               `bson:"host" json:"host"`
	Location    string             `bson:"location" json:"location"`
	SeatCount   int                `bson:"seatcount" json:"seatcount"`
	Seats       []Seat             `bson:"seats" json:"seats"`
	Bookings    []Booking          `bson:"bookings" json:"bookings"`
}
