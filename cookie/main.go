package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.Header.Get("Cookie"))
		//writer.Header().Add("set-cookie", "xxx=xxxxxx")
		sid := &http.Cookie{
			Name:       "sid",
			Value:      "xxxx",
			Path:       "",
			Domain:     "",
			Expires:    time.Time{},
			RawExpires: "",
			MaxAge:     0,
			Secure:     false,
			HttpOnly:   false,
			SameSite:   0,
			Raw:        "",
			Unparsed:   nil,
		}
		http.SetCookie(writer, sid)
	})
	http.ListenAndServe("0.0.0.0:9999", nil)
}
