package main

import (
	"fmt"
	"ticketbook/config"
	"ticketbook/routers"
)

func main() {
	fmt.Println(" ticket booking app ")
	config.DBConnect()
	routers.Router()

}
