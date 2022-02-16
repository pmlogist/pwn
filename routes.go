package main

import (
	"database/sql"
	"log"
	"net/http"
	"net/url"
	"strings"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

type Router struct {
	db        *sql.DB
	CsrfToken string
}

type M map[string]interface{}

func (router *Router) NiceTry(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, _ := template.ParseFiles("./templates/nice-try.html")

	router.SetCsrf(w, r)
	t.Execute(w, nil)
}

func (router *Router) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	home, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Print(err)
	}

	router.ClearSession(w, r)
	router.SetCsrf(w, r)
	home.Execute(w, nil)
}

func (router *Router) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, _ := template.ParseFiles("./templates/index.html")

	username := r.FormValue("username")
	password := r.FormValue("password")
	user, _ := router.GetOneByUsername(username)

	if username == "" || password == "" {
		t.Execute(w, nil)
		router.SetCsrf(w, r)
		return
	}

	if username != user.Username || password != user.Password {
		t.Execute(w, nil)
		router.SetCsrf(w, r)
	} else {
		w.Header().Set("Set-Cookie", "uid="+user.Id)
		router.SetCsrf(w, r)
		http.Redirect(w, r, "/profile"+"?user_id="+user.Id, http.StatusSeeOther)
	}

}

func (router *Router) Profile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	uid, err := r.Cookie("uid")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		router.SetCsrf(w, r)
		return
	}
	if r.FormValue("logout") != "" {
		c := &http.Cookie{
			Name:     "x-csrf",
			Value:    "",
			Path:     "/",
			HttpOnly: true,
			MaxAge:   0,
		}

		http.SetCookie(w, c)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		router.SetCsrf(w, r)
		return
	}

	// Safe
	//userId := r.URL.Query().Get("user_id")

	users, _ := router.GetAll()

	userId := strings.Split(r.URL.RawQuery, "user_id=")[1]
	decodedUserId, _ := url.QueryUnescape(userId)
	user, _ := router.GetOneById(decodedUserId)
	user.Id = decodedUserId

	if uid.Value != user.Id && user.IsAdmin == true {
		http.Redirect(w, r, "/profile?user_id="+uid.Value, http.StatusSeeOther)
		router.SetCsrf(w, r)
		return
	}

	t, _ := template.ParseFiles("./templates/profile.html")
	router.SetCsrf(w, r)
	t.Execute(w, M{"user": user, "users": users})
}

func (router *Router) ProfileLikesAndDelete(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//t, _ := template.ParseFiles("./templates/profile-deleted.html")
	likesProfileId := r.FormValue("likes")
	deleteProfileId := r.FormValue("delete")
	router.CheckCsrf(w, r)

	if likesProfileId != "" {
		router.LikeOneById(likesProfileId)
		queryUserId := r.URL.Query().Get("user_id")
		http.Redirect(w, r, "/profile"+"?user_id="+queryUserId, http.StatusSeeOther)

	}

	if deleteProfileId != "" {
		router.DeleteOneById(deleteProfileId)
		uid, _ := r.Cookie("uid")
		http.Redirect(w, r, "/profile"+"?user_id="+uid.Value, http.StatusSeeOther)
	}

}
