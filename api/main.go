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
	// user handler
	router.POST("/user/", handler.UserCreate)
	router.POST("/user/:username", handler.Login)
	router.GET("/user/status/", handler.IsLogin)
	router.GET("/user/logout/", handler.Logout)
	router.DELETE("/user", handler.DeleteUser)
	// video handler
	router.POST("/video", handler.VideoCreate)
	router.GET("/video/:video_id", handler.VideoGet)
	router.DELETE("/video", handler.VideoCreate)
	// comment handler
	router.POST("/comment", handler.CommentCreate)
	router.GET("/comments/:video_id", handler.GetComments)

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

	r := RegisterHandlers()
	http.ListenAndServe(":8080", r)
}
