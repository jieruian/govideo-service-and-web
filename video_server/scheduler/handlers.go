package main

import (
	"github.com/julienschmidt/httprouter"
	"govideo/video_server/scheduler/dbops"
	"net/http"
)

func vidDelRecHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	vid := p.ByName("vid-id")
	if len(vid) == 0 {
		sendErrorResponse(w, 400, "Video id should not be empty")
		return
	}
	err := dbops.AddDeletionRecode(vid)
	if err != nil {
		sendErrorResponse(w, 500, "Internal server error")
		return
	}
	sendNormalResponse(w, 200, "删除成功啦😜😜")
	return
}
