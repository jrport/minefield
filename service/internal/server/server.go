package server

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"jrport/minefield/internal/pool"
	"net/http"
)

type GameServer struct {
	port           string
	Mux            *http.ServeMux
	maxPlayerCount int
}

func NewGameServer(port string, limit int) *GameServer {
	mux := http.NewServeMux()

	return &GameServer{
		port:           port,
		Mux:            mux,
		maxPlayerCount: limit,
	}
}

func SetupRoutes(muxer *http.ServeMux, gp *pool.MatchPool) {
	rc := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		},
	)

	uidHandler := newHandleWithRedis(rc, getAnomUserId)
	matchMakerHandler := newMatchMakerHandler(gp, rc)

	muxer.Handle("POST /anom_match", uidHandler)
	muxer.Handle("GET /join_match", matchMakerHandler)
}

func (gs *GameServer) Run() error {
	gp := pool.NewMatchPool(gs.maxPlayerCount)
	SetupRoutes(gs.Mux, gp)
	sr := make(chan error)

	fmt.Println("Running server")

	go func() {
		if err := http.ListenAndServe(gs.port, gs.Mux); err != nil {
			sr <- err
		}
	}()

	return <-sr
}
