package main

import (
	"log"
	"net/http"

	"github.com/erhemdiputra/go-crud/database"
	"github.com/erhemdiputra/go-crud/user/delivery"
	"github.com/erhemdiputra/go-crud/views"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func init() {
	if err := views.Init(); err != nil {
		panic(err)
	}

	if err := database.Init(); err != nil {
		panic(err)
	}
}

func main() {
	router := httprouter.New()

	delivery.NewUserHandler(router)

	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", router)
}
