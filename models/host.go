package models

type Host struct {
	Name        string   `bson:"name" json:"name"`
	Description string   `bson:"description" json:"description"`
	Images      []string `bson:"images" json:"images"`
	Rating      int      `bson:"rating" json:"rating"`
	Viewers     int      `bson:"viewers" json:"viewers"`
	ShowsCount  int      `bson:"showscount" json:"showscount"`
	FeePerHour  int      `bson:"feeperhour" json:"feeperhour"`
	BaseFee     int      `bson:"basefee" json:"basefee"`
}
