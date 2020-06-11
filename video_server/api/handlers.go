package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"govideo/video_server/api/dbops"
	"govideo/video_server/api/defs"
	"govideo/video_server/api/session"
	"govideo/video_server/api/utils"
	"io/ioutil"
	"log"
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
	res, _ := ioutil.ReadAll(r.Body)
	log.Println(res)
	ubody := &defs.UserCredential{}
	if err := json.Unmarshal(res, ubody); err != nil {
		log.Println(err)
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	userName := p.ByName("user_name")
	log.Printf("Login url name: %s", userName)
	log.Printf("Login body name: %s", ubody.Username)
	if userName != ubody.Username {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}
	pwd, err := dbops.GetUserCredential(userName)
	log.Printf("Login pwd: %s", pwd)
	if err != nil || len(pwd) == 0 || pwd != ubody.Pwd {
		si := defs.SignedUp{
			Success:   false,
			SessionId: "",
			Message:   "用户不存在",
		}
		resp, _ := json.Marshal(si)
		sendNormalResponse(w, string(resp), 400)
		return
	}

	id := session.GenerateNewSessionId(userName)
	si := defs.SignedUp{
		Success:   true,
		SessionId: id,
		Message:   "恭喜！登录成功啦",
	}
	if resp, err := json.Marshal(si); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}
}

func CreateUserInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if !validateUser(w, r) {
		log.Printf("Unauthorized user \n")
	}
	userName := ps.ByName("username")
	u, err := dbops.GetUser(userName)
	if err != nil {
		log.Printf("Erorr in GetUserinfo: %s", err)
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	ui := &defs.UserInfo{Id: u.Id}
	if resp, err := json.Marshal(ui); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}

}

func AddNewVideo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if !validateUser(w, r) {
		log.Printf("Unauthorized user \n")
		return
	}
	res, _ := ioutil.ReadAll(r.Body)
	nvbody := &defs.NewVideo{}
	if err := json.Unmarshal(res, nvbody); err != nil {
		log.Printf("%s", err)
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	vi, err := dbops.AddNewVideo(nvbody.AuthorId, nvbody.Name)
	if err != nil {
		log.Printf("Error in AddNewVideo: 5s", err)
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	if resp, err := json.Marshal(vi); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 201)
	}
}

func ListAllVideos(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if !validateUser(w, r) {
		log.Printf("Unauthorized user \n")
		return
	}
	uname := ps.ByName("username")
	vs, err := dbops.ListVideoInfo(uname, 0, utils.GetCurrentTimestampSec())
	if err != nil {
		log.Printf("Error in ListAllVideos: %s", err)
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	vsi := &defs.VideosInfo{Videos: vs}

	if resp, err := json.Marshal(vsi); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}

}

func DeleteVideo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if !validateUser(w, r) {
		log.Printf("Unauthorized user \n")
		return
	}

	vid := ps.ByName("vid-id")
	err := dbops.DeleteVideoInfo(vid)
	if err != nil {
		log.Printf("Error in DeleteVideo: %s", err)
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	sendNormalResponse(w, "删除视频成功", 204)
}

func PostComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if !validateUser(w, r) {
		log.Printf("Unauthorized user \n")
		return
	}

	reqBody, _ := ioutil.ReadAll(r.Body)
	cbody := &defs.NewComment{}
	if err := json.Unmarshal(reqBody, cbody); err != nil {
		log.Printf("%s", err)
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	vid := ps.ByName("vid-id")
	if err := dbops.AddNewComments(vid, cbody.AuthorId, cbody.Content); err != nil {
		log.Printf("Error in PostComment: %s", err)
		sendErrorResponse(w, defs.ErrorDBError)
	} else {
		sendNormalResponse(w, "ok", 201)
	}

}

func ShowComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if !validateUser(w, r) {
		log.Printf("Unauthorized user \n")
		return
	}
	vid := ps.ByName("vid-id")
	cm, err := dbops.ListComments(vid, 0, utils.GetCurrentTimestampSec())
	if err != nil {
		log.Printf("Error in ShowComments: %s", err)
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	cms := &defs.Comments{Comments: cm}
	if resp, err := json.Marshal(cms); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}
}
