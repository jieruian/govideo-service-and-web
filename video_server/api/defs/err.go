package defs

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSC int
	Error  Err
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
	ErrorDBError = ErrorResponse{
		HttpSC: 500, //内部错误
		Error: Err{
			Error:     "DB ops faild",
			ErrorCode: "003",
		},
	}
	ErrorInternalFaults = ErrorResponse{
		HttpSC: 500,
		Error: Err{
			Error:     "Internal Server Fail",
			ErrorCode: "004",
		},
	}
)
