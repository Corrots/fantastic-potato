package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/corrots/fantastic-potato/api/response"

	"github.com/corrots/fantastic-potato/api/database"
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
	//body, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	//io.WriteString(w, err.Message())
	//	response.Message(w, &response.ErrorRequestBodyParseFailed)
	//	return
	//}
	if err := parseParameter(r.Body, &user); err != nil {
		response.Error(w, &response.ErrorParseParameter)
		return
	}
	fmt.Printf("user: %+v\n", user)
	if !isPasswordConfirmed(&user) {
		response.Error(w, &response.ErrorPasswordConfirmed)
		return
	}
	if err := user.Add(); err != nil {
		response.Error(w, &response.ErrorUserAddFailed)
		return
	}
	//io.WriteString(w, "created")
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
		response.Error(w, &response.ErrorParseParameter)
		return
	}
	err := user.GetByUsername()
	if err != nil {
		response.Error(w, &response.ErrorDBError)
		return
	}
	fmt.Println(user.Password)
}
