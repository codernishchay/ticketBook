package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ticketbook/models"
)

func Tickerhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ticket booking start")
	var ticket models.Booking
	err := json.NewDecoder(r.Body).Decode(&ticket)
	if err != nil {
		fmt.Println(err)
	}
	// if seat is avilable ?  book else return
	json.NewEncoder(w).Encode(ticket)
}
