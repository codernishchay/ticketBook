package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"ticketbook/config"
	"ticketbook/models"
	"time"
)

// CreateNewShow : create shows in database
func CreateNewShow(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create a show")
	var f models.Show
	err := json.NewDecoder(r.Body).Decode(&f)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	res, er := config.DB.Collection("shows").InsertOne(ctx, f)
	if er != nil {
		fmt.Println(er)
	}
	json.NewEncoder(w).Encode(res)
}
