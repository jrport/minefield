package main

import (
	"fmt"
	"net/http"
	"strings"
)

func (u *UserManager) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ip := strings.Split(r.RemoteAddr, ":")[0]
	r.ParseForm()

	name := strings.TrimSpace(r.Form.Get("name"))
	if name == "" {
		http.Error(w, "Missing name field in form.", 422)
		println("Bad user form rejected request.")
		return
	}

	user := NewUser(name, ip)
	if ok := u.QueueUser(user); !ok {
		http.Error(w, "Full server.", 422)
		fmt.Printf(
			"Full server. Rejected one request.\nCurrent User Count: %v\nMax supported: %v",
			u.UserCount, u.MaxCapacity,
		)
		return
	}

	u.PrintStats()
	fmt.Fprint(w, user.Id)
	w.WriteHeader(201)
}

func main() {
	muxer := http.NewServeMux()
	userRegister := NewUserRegister(1024)
	userMatchMaker := NewUserMatchMaker(1024)

	go MatchMakingService(userMatchMaker, userRegister)

	muxer.Handle("POST /anom_session", userRegister)
	muxer.Handle("POST /match", userMatchMaker)

	http.ListenAndServe(":8080", muxer)
}
