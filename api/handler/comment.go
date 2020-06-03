package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/corrots/fantastic-potato/api/tools/session"

	"github.com/corrots/fantastic-potato/api/response"
	"github.com/corrots/fantastic-potato/api/tools/database"
	"github.com/julienschmidt/httprouter"
)

func CommentCreate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// check if user logged in
	user, err := session.LoginInfo(r)
	if err != nil {
		response.Error(w, &response.GetLoginInfoErr, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, &response.ParseRequestErr, err)
		return
	}
	defer r.Body.Close()
	var c database.Comment
	if err := json.Unmarshal(body, &c); err != nil {
		response.Error(w, &response.JsonDecodeErr, err)
		return
	}
	c.UserId = user.UserId
	c.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	if err := c.Add(); err != nil {
		response.Error(w, &response.DbError, err)
		return
	}
	response.OK(w, http.StatusOK, map[string]interface{}{
		"code":    response.SucceedCode,
		"message": "comment created",
	})
}

func GetComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	videoId, _ := strconv.Atoi(ps.ByName("video_id"))
	comments, err := database.GetComments(videoId)
	if err != nil {
		response.Error(w, &response.DbError, err)
		return
	}
	response.OK(w, http.StatusOK, map[string]interface{}{
		"code":    response.SucceedCode,
		"message": "succeed",
		"data":    comments,
	})
}
