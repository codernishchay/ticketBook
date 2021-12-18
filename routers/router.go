package routers

import (
	"fmt"
	"net/http"
	"ticketbook/handler"

	"github.com/gorilla/mux"
)

func Router() {
	r := mux.NewRouter()
	r.HandleFunc("/create", handler.CreateNewShow)
	r.HandleFunc("/book", handler.Tickerhandler)

	fmt.Println("Router is here")
	http.ListenAndServe(":8080", r)
}
