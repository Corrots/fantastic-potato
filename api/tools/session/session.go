package session

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/corrots/fantastic-potato/api/tools/database"
)

const (
	expiration = 86400
	sessionLen = 32
	sessionKey = "fantastic-potato"
)

func Login(w http.ResponseWriter, u *database.User) error {
	sessionID, err := newSessionID()
	if err != nil {
		return err
	}
	value, _ := json.Marshal(u)
	if err := Set(sessionID, value, expiration*time.Second); err != nil {
		return err
	}
	expire := time.Now().Add(time.Second * expiration)
	http.SetCookie(w, &http.Cookie{
		Name:     sessionKey,
		Value:    sessionID,
		Path:     "/",
		Expires:  expire,
		HttpOnly: true,
	})
	return nil
}

func LoginInfo(r *http.Request) (*database.User, error) {
	cookie, err := r.Cookie(sessionKey)
	if err != nil {
		return nil, fmt.Errorf("get session id from cookie failed: %v\n", err)
	}
	var value string
	value, err = Get(cookie.Value)
	if err != nil {
		return nil, err
	}
	u := new(database.User)
	json.Unmarshal([]byte(value), u)
	return u, nil
}

func Logout(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie(sessionKey)
	if err != nil {
		return fmt.Errorf("get session id from cookie failed: %v\n", err)
	}
	if err := Set(cookie.Value, nil, -1); err != nil {
		return fmt.Errorf("remove cookie failed")
	}
	expire := time.Now().Add(time.Second * -1)
	http.SetCookie(w, &http.Cookie{
		Name:     sessionKey,
		Value:    "",
		Path:     "/",
		Expires:  expire,
		HttpOnly: true,
	})
	return nil
}

func newSessionID() (string, error) {
	b := make([]byte, sessionLen)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), err
}
