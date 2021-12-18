package main

import (
	"fmt"
	"ticketbook/db"
	"ticketbook/routers"
)

func main() {
	fmt.Println(" ticket booking app ")
	db.DBConnect()
	routers.Router()

}
