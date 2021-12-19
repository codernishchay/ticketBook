package handler

import (
	"context"
	"ticketbook/config"
	"ticketbook/models"

	"go.mongodb.org/mongo-driver/bson"
)

var SeatCheck []models.Seat

// CheckAvilability will return number of seats
func CheckAvilability(ticket models.Booking) []models.Seat {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	showID := ticket.ShowID
	show := config.DB.Collection("shows").FindOne(ctx, bson.M{"_id": showID})
	var sho models.Show
	show.Decode(sho)
	return sho.Seats
}

// CheckSeatStatus will return if that seat is avilable or not?
func CheckSeatStatus(seatNo int) bool {
	for i := 0; i < len(SeatCheck); i++ {
		if seatNo == SeatCheck[i].Number {
			if SeatCheck[i].Booked == true {
				return false
			}
			return true
		}
	}
	return true
}
