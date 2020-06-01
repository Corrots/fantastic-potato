package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/corrots/fantastic-potato/api/tools/session"

	"github.com/corrots/fantastic-potato/api/response"
	"github.com/corrots/fantastic-potato/api/tools/database"
	"github.com/julienschmidt/httprouter"
)

func parseParameter(b io.ReadCloser, u *database.User) error {
	body, err := ioutil.ReadAll(b)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, &u)
}

func UserCreate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var user database.User
	if err := parseParameter(r.Body, &user); err != nil {
		response.Error(w, &response.JsonDecodeErr, err)
		return
	}
	fmt.Printf("user: %+v\n", user)
	if !isPasswordConfirmed(&user) {
		response.Error(w, &response.ConfirmedErr, fmt.Errorf("password confirmed failed"))
		return
	}
	if err := user.Add(); err != nil {
		response.Error(w, &response.UserCreateErr, err)
		return
	}
	response.OK(w, http.StatusCreated, map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "created",
	})
}

func isPasswordConfirmed(u *database.User) bool {
	return u.Password == u.ConfirmPassword
}

func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var user database.User
	if err := parseParameter(r.Body, &user); err != nil {
		response.Error(w, &response.JsonDecodeErr, err)
		return
	}
	user.Username = ps.ByName("username")
	//user.Password = ps.ByName("password")
	//user.ConfirmPassword = ps.ByName("confirm_password")
	output, err := user.GetByUsername()
	if err != nil {
		response.Error(w, &response.DbError, err)
		return
	}
	if user.Password != output.Password {
		response.Error(w, &response.UserAuthErr, err)
		return
	}
	if err = session.Login(w, output); err != nil {
		response.Error(w, &response.UserLoginErr, fmt.Errorf("user login err: %v\n", err))
		return
	}

	response.OK(w, http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "login successful",
	})
}

func IsLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	u, err := session.LoginInfo(r)
	if err != nil {
		response.Error(w, &response.GetLoginInfoErr, fmt.Errorf("get login info err: %v\n", err))
		return
	}
	//io.WriteString(w, fmt.Sprintf("%+v\n", u))
	response.OK(w, http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "is login",
		"data":    u,
	})
}

func Logout(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if err := session.Logout(w, r); err != nil {
		log.Println(err)
		return
	}
	response.OK(w, http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "logout successful",
	})
}
