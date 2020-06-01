package main

import (
	"log"
	"net/http"

	"github.com/corrots/fantastic-potato/api/tools/session"

	"github.com/corrots/fantastic-potato/api/config"
	"github.com/corrots/fantastic-potato/api/handler"
	"github.com/corrots/fantastic-potato/api/tools/database"
	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user/", handler.UserCreate)
	router.POST("/user/:username", handler.Login)
	//router.GET("/user/login/:username/:password/:confirm_password", handler.Login)
	router.GET("/user/status/", handler.IsLogin)
	router.GET("/user/logout/", handler.Logout)
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
	session.Init()
	//session.IsLogin()

	r := RegisterHandlers()
	http.ListenAndServe(":8080", r)
}
