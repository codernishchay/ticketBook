package models

type Venue struct {
	NearBy      string `bson:"nearby" json:"nearby"`
	City        string `bson:"city" json:"city"`
	Seats       []Seat `bson:"seats" json:"seats"`
	Name        string `bson:"name" json:"name"`
	Email       string `bson:"email" json:"email"`
	PhoneNumber string `bson:"phonenumber" json:"phonenumber"`
	Manager     string `bson:"manager" json:"manager"`
	FeePerHour  int    `bson:"feeperhour" json:"feeperhour"`
	BaseFee     int    `bson:"basefee" json:"basefee"`
}
