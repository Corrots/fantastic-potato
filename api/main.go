package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/corrots/fantastic-potato/api/config"
	"github.com/corrots/fantastic-potato/api/database"
	"github.com/corrots/fantastic-potato/api/handler"
	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user/", handler.UserCreate)
	router.POST("/user/:username", handler.Login)
	return router
}

func main() {
	if err := config.Init(); err != nil {
		log.Printf("load config err: %v\n", err)
		panic(err)
	}
	if err := database.Init(); err != nil {
		log.Println(err)
		panic(err)
	}

	r := RegisterHandlers()
	fmt.Println(http.ListenAndServe(":8080", r))
}
