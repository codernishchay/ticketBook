package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"ticketbook/config"
	"ticketbook/models"

	"go.mongodb.org/mongo-driver/mongo"
)

//considiring the final seat for booking
// Tickethandler will book ticket for same seat.

func Tickerhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ticket booking start")
	var ticket models.Booking

	err := json.NewDecoder(r.Body).Decode(&ticket)
	if err != nil {
		fmt.Println(err)
	}
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	var result *mongo.InsertOneResult
	// verify if seat is free
	if CheckSeatStatus(ticket.Seat.Number) {
		result, err = config.DB.Collection("bookings").InsertOne(ctx, ticket)
		if err != nil {
			log.Fatal("Error , Cant book ticket")
		}
	} else {
		json.NewEncoder(w).Encode("some  error occured ")
	}
	json.NewEncoder(w).Encode(result)
}
