package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
)

type HomePage struct {
	Name string
}
type UserPage struct {
	Name string
}

func homeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	cname, err := r.Cookie("username")
	sessionId, err2 := r.Cookie("session")
	if err != nil || err2 != nil {
		p := &HomePage{Name: "lsy"}
		t, err := template.ParseFiles("./video_server/templates/home.html")
		if err != nil {
			log.Printf("Parsing template home.html error: %s", err)
			return
		}
		t.Execute(w, p)
		return
	}

	if len(sessionId.Value) != 0 && len(cname.Value) != 0 {
		//重定向
		http.Redirect(w, r, "/userhome", http.StatusFound)
		return
	}
}

func userhomeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cname, err := r.Cookie("username")
	_, err2 := r.Cookie("session")
	if err != nil || err2 != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	fname := r.FormValue("username")
	var p *UserPage
	if len(cname.Value) != 0 {
		p = &UserPage{Name: cname.Value}
	} else if len(fname) != 0 {
		p = &UserPage{Name: fname}
	}

	t, e := template.ParseFiles("./templates/userhome.html")
	if e != nil {
		log.Printf("Parsing userhome.html error: %s", e)
		return
	}
	t.Execute(w, p)

}