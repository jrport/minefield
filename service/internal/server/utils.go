package server

import (
	"context"
	"net/http"

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

var ctx = context.Background() // TODO redis timeout error, make custom context with timeout and custom error type
