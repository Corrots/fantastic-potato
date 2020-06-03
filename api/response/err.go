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
	SucceedCode = "00"

	ParseRequestErr = ErrorResponse{StatusCode: 400, JsonResp: Resp{Code: "-1", Message: "Request body is not correct"}}
	JsonDecodeErr   = ErrorResponse{StatusCode: 402, JsonResp: Resp{Code: "-2", Message: "Decode request json body failed"}}
	DbError         = ErrorResponse{StatusCode: 500, JsonResp: Resp{Code: "-3", Message: "DB ops failed"}}
	ServerErr       = ErrorResponse{StatusCode: 500, JsonResp: Resp{Code: "-4", Message: "Internal service error"}}

	ConfirmedErr    = ErrorResponse{StatusCode: 402, JsonResp: Resp{Code: "11001", Message: "Password confirmed failed"}}
	UserAuthErr     = ErrorResponse{StatusCode: 401, JsonResp: Resp{Code: "11002", Message: "User authentication failed"}}
	UserCreateErr   = ErrorResponse{StatusCode: 500, JsonResp: Resp{Code: "11003", Message: "User add failed"}}
	UserLoginErr    = ErrorResponse{StatusCode: 500, JsonResp: Resp{Code: "11004", Message: "User login failed"}}
	GetLoginInfoErr = ErrorResponse{StatusCode: 500, JsonResp: Resp{Code: "11005", Message: "Get login info failed"}}
)
