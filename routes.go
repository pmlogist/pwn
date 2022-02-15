package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

type Router struct {
	db *sql.DB
}

type M map[string]interface{}

var csrf int

func CsrfToken() string {
	csrf = rand.Intn(100000)
	return fmt.Sprintf("x-csrf=%v", csrf)
}

func (router *Router) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, _ := template.ParseFiles("./templates/index.html")

	t.Execute(w, nil)
}

func (router *Router) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, _ := template.ParseFiles("./templates/index.html")

	username := r.FormValue("username")
	password := r.FormValue("password")
	user, _ := router.GetOneByUsername(username)

	if username == "" || password == "" {
		t.Execute(w, nil)
		return
	}

	if username != user.Username || password != user.Password {
		t.Execute(w, nil)
	} else {
		w.Header().Set("Set-Cookie", CsrfToken())
		w.Header().Set("Set-Cookie", "uid="+user.Id)
		http.Redirect(w, r, "/profile"+"?user_id="+user.Id, http.StatusSeeOther)
	}

}

func (router *Router) Profile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//csrfCookie, err := r.Cookie("x-csrf")
	uid, err := r.Cookie("uid")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	//if csrfCookie.Value != fmt.Sprint(csrf) {
	//	http.Redirect(w, r, "/", http.StatusSeeOther)
	//	return
	//}

	// Safe
	userId := r.URL.Query().Get("user_id")

	//userId := strings.Split(r.URL.RawQuery, "user_id=")[1]
	//decodedUserId, _ := url.QueryUnescape(userId)
	user, _ := router.GetOneById(userId)
	user.Id = userId

	if uid.Value != user.Id && user.IsAdmin == true {
		http.Redirect(w, r, "/profile?user_id="+uid.Value, http.StatusSeeOther)
		return
	}

	users, _ := router.GetAll()
	t, _ := template.ParseFiles("./templates/profile.html")
	t.Execute(w, M{"user": user, "users": users})
}

func (router *Router) ProfileLikesAndDelete(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//t, _ := template.ParseFiles("./templates/profile-deleted.html")
	likes := r.FormValue("likes")

	if likes != "" {
		router.LikeOneById(likes)
		queryUserId := r.URL.Query().Get("user_id")
		http.Redirect(w, r, "/profile"+"?user_id="+queryUserId, http.StatusSeeOther)
	}

	//csrfCookie, err := r.Cookie("x-csrf")
	//if err != nil {
	//	http.Redirect(w, r, "/", http.StatusSeeOther)
	//	return
	//}

	//if csrfCookie.Value != fmt.Sprint(csrf) {
	//	http.Redirect(w, r, "/", http.StatusSeeOther)
	//	return
	//}

	//router.DeleteOneById(r.FormValue("id"))

	//t.Execute(w, nil)

}
