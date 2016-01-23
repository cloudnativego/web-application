package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/unrolled/render"
)

type sampleContent struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

func testHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, sampleContent{ID: "8675309", Content: "Hello from Go!"})
	}
}

func cookieWriteHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "sample", Value: "this is a cookie", Expires: expiration}
		http.SetCookie(w, &cookie)
		formatter.JSON(w, http.StatusOK, "cookie set")
	}
}

func cookieReadHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie("sample")
		if err == nil {
			fmt.Fprint(w, cookie.Value)
		} else {
			fmt.Fprintf(w, "failed to read cookie, %v", err)
		}
	}
}
