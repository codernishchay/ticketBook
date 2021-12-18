package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"ticketbook/db"
	"ticketbook/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Tickerhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ticket booking start")
	var ticket models.Booking
	err := json.NewDecoder(r.Body).Decode(&ticket)
	// search for the show
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	showID := ticket.ShowID
	show := db.DB.Collection("shows").FindOne(ctx, bson.M{"_id": showID})
	fmt.Println(show)
	var sho models.Show
	show.Decode(sho)
	var insertResult *mongo.InsertOneResult
	if sho.SeatCount >= 0 {
		// confirm booking
		fmt.Println("money paid, well done ")
		// pay money
		insertResult, err = db.DB.Collection("bookings").InsertOne(ctx, ticket)
		if err != nil {
			fmt.Println("cant book your ticket, retry ")
		}

		// check payment again and then
		// send notification and update user database

	}
	// check if there are seats left ?
	// if seats are left then check if requierd seat is left
	// confirm booking
	// pay money
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(insertResult)
}
