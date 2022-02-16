package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func RandSeq(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}

func (router *Router) SetCsrf(w http.ResponseWriter, r *http.Request) {
	randomString := RandSeq(10)
	c := &http.Cookie{
		Name:  "x-csrf",
		Value: randomString,
		Path:  "/",
		//HttpOnly: true,
	}

	http.SetCookie(w, c)
	router.CsrfToken = randomString
}

func (router *Router) ClearSession(w http.ResponseWriter, r *http.Request) {
	randomString := RandSeq(10)
	c := &http.Cookie{
		Name:   "x-csrf",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
		//HttpOnly: true,
	}

	c2 := &http.Cookie{
		Name:   "uid",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
		//HttpOnly: true,
	}

	http.SetCookie(w, c)
	http.SetCookie(w, c2)
	router.CsrfToken = randomString
}

func (router *Router) CheckCsrf(w http.ResponseWriter, r *http.Request) error {
	token, err := r.Cookie("x-csrf")
	if err != nil {
		http.Redirect(w, r, "/nice-try", http.StatusSeeOther)
		return nil
	}
	if token.Value != router.CsrfToken {
		http.Redirect(w, r, "/nice-try", http.StatusSeeOther)
		return err
	} else {
		return nil
	}
}
