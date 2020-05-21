package main

import (
	"encoding/json"
	"govideo/video_server/api/defs"
	"io"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, errResponse defs.ErrorResponse) {
	w.WriteHeader(errResponse.HttpSC)
	resStr, _ := json.Marshal(&errResponse.Error)
	io.WriteString(w, string(resStr))
}

func sendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
