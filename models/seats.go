package models

type Seat struct {
	Number int  `bson:"number" json:"number"`
	Val    int  `bson:"value" json:"value"`
	Booked bool `bson:"booked" json:"booked"`
}
