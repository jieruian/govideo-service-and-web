package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"time"
)

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	vl := VIDEO_OIR + vid + ".mp4"
	//tmp := "/Users/lishaoyu/go/src/govideo/video_server/videos/testvideo.mp4"
	video, err := os.Open(vl)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "internal error")
		return
	}
	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)
	defer video.Close()

}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
