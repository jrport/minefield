package server

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type handleWithRedis struct {
	client     *redis.Client
	handleFunc func(http.ResponseWriter, *http.Request, *redis.Client)
}

func (hr *handleWithRedis) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hr.handleFunc(w, r, hr.client)
}

func newHandleWithRedis(rc *redis.Client, hf func(http.ResponseWriter, *http.Request, *redis.Client)) *handleWithRedis {
	return &handleWithRedis{
		client:     rc,
		handleFunc: hf,
	}
}

var ctx = context.Background() // TODO redis  timeout error, make custom context with timeout and custom error type

func getUserId(w http.ResponseWriter, r *http.Request, rc *redis.Client) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Form", 422)
		return
	}

	name := strings.TrimSpace(r.PostForm.Get("name"))
	if len(name) == 0 {
		http.Error(w, "Bad Name", 422)
		return
	}

	ip := []byte(r.RemoteAddr)
	fmt.Printf("incoming request from: %v\n", ip)

	dst := make([]byte, base64.StdEncoding.EncodedLen(len(ip)))
	base64.StdEncoding.Encode(dst, ip)

	fmt.Printf("ip: %s\nb64(key): %s\nname(val): %v\n", ip, string(dst), name)
	err = rc.Set(ctx, string(dst), name, time.Minute*2).Err()
	if err != nil {
		fmt.Println("Redis error: ", err.Error())
		http.Error(w, "Internal Error", 500)
	}

	_, err = fmt.Fprint(w, string(dst))
	if err != nil {
		fmt.Println("Error on write")
		http.Error(w, "Internal Error", 500)
	}
}

func SetupRoutes(muxer *http.ServeMux) {
	rc := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		},
	)

	uidHandle := newHandleWithRedis(rc, getUserId)

	muxer.Handle("POST /anom_match", uidHandle)
}
