package response

type Resp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type ErrorResponse struct {
	StatusCode int
	JsonResp   Resp
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{StatusCode: 400, JsonResp: Resp{Code: "-1", Message: "Request body is not correct"}}
	ErrorParseParameter         = ErrorResponse{StatusCode: 402, JsonResp: Resp{Code: "-2", Message: "Decode request json body failed"}}
	ErrorDBError                = ErrorResponse{StatusCode: 500, JsonResp: Resp{Code: "-3", Message: "DB ops failed"}}
	ErrorInternalFaults         = ErrorResponse{StatusCode: 500, JsonResp: Resp{Code: "-4", Message: "Internal service error"}}

	ErrorPasswordConfirmed = ErrorResponse{StatusCode: 402, JsonResp: Resp{Code: "11001", Message: "Password confirmed failed"}}
	ErrorNotAuthUser       = ErrorResponse{StatusCode: 401, JsonResp: Resp{Code: "11002", Message: "User authentication failed."}}
	ErrorUserAddFailed     = ErrorResponse{StatusCode: 500, JsonResp: Resp{Code: "11003", Message: "User add failed."}}
)
