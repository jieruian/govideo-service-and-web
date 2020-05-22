package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
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
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "文件太大")
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
	}
	fileName := p.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_OIR+fileName+".mp4", data, 0666)
	if err != nil {
		log.Printf("Write file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Upload successfully")
}

func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, _ := template.ParseFiles("/Users/lishaoyu/go/src/govideo/video_server/streamserve/upload.html")
	t.Execute(w, nil)
}
