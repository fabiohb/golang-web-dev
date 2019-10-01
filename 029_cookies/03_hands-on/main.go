package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	cookie, cookieErr := req.Cookie("my-cookie")
	if cookieErr != nil {
		cookie = &http.Cookie{
			Name:  "my-cookie",
			Value: "1",
		}
	} else {
		count, err := strconv.Atoi(cookie.Value)
		if err != nil {
			count = 0
		}
		count++
		cookie.Value = strconv.Itoa(count)
	}

	http.SetCookie(w, cookie)

	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}
