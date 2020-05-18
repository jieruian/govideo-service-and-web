package defs

type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSC int
	Error Err
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{
		HttpSC: 400,
		Error: Err{
			Error:     "request body isn't correct",
			ErrorCode: "002",
		},
	}
	ErrorNotAuthUser = ErrorResponse{
		HttpSC: 401,
		Error: Err{
			Error:     "user authentication failed",
			ErrorCode: "002",
		},
	}

)