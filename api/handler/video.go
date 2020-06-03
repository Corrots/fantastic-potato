package handler

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/corrots/fantastic-potato/api/response"
	"github.com/corrots/fantastic-potato/api/tools/database"
	"github.com/corrots/fantastic-potato/api/tools/session"
	"github.com/julienschmidt/httprouter"
)

func VideoCreate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	u, err := session.LoginInfo(r)
	if err != nil {
		io.WriteString(w, "not login")
		return
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, &response.ParseRequestErr, err)
		return
	}
	defer r.Body.Close()
	var v database.Video
	if err := json.Unmarshal(b, &v); err != nil {
		response.Error(w, &response.JsonDecodeErr, err)
		return
	}
	v.AuthorId = u.UserId
	v.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	if err := v.Add(); err != nil {
		response.Error(w, &response.ServerErr, err)
		return
	}
	response.OK(w, http.StatusCreated, map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "created",
	})
}

func VideoGet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var v database.Video
	videoId, _ := strconv.Atoi(ps.ByName("video_id"))
	v.VideoId = videoId
	video, err := database.GetVideoById(videoId)
	if err != nil {
		response.Error(w, &response.ServerErr, err)
		return
	}
	response.OK(w, http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "succeed",
		"data":    video,
	})
}
