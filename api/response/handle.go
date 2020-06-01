package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func OK(w http.ResponseWriter, statusCode int, resp map[string]interface{}) {
	w.WriteHeader(statusCode)
	b, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	w.Write(b)
}

func Error(w http.ResponseWriter, resp *ErrorResponse, err error) {
	w.WriteHeader(resp.StatusCode)
	res, _ := json.Marshal(&resp.JsonResp)
	log.Println(err)
	w.Write(res)
}
