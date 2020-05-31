package response

import (
	"encoding/json"
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

func Error(w http.ResponseWriter, resp *ErrorResponse) {
	w.WriteHeader(resp.StatusCode)
	res, err := json.Marshal(&resp.JsonResp)
	if err != nil {
		panic(err)
	}
	w.Write(res)
}
