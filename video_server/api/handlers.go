package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"govideo/video_server/api/dbops"
	"govideo/video_server/api/defs"
	"govideo/video_server/api/session"
	"io"
	"io/ioutil"
	"net/http"
)

func RegistUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	fmt.Println(res)
	uBody := &defs.UserCredential{}
	//res = bytes.TrimPrefix(res, []byte("\xef\xbb\xbf"))
	if err := json.Unmarshal(res, uBody); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	if err := dbops.AddUserCredential(uBody.Username, uBody.Pwd); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	id := session.GenerateNewSessionId(uBody.Username)
	su := defs.SignedUp{
		Success:   true,
		SessionId: id,
	}
	if res, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(res), 200)
	}

}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userName := p.ByName("user_name")
	io.WriteString(w, userName)
}
