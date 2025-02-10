package server

import (
	"fmt"
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

func (gs *GameServer) Run() error {
	SetupRoutes(gs.Mux)

	fmt.Println("Running server")
	if err := http.ListenAndServe(gs.port, gs.Mux); err != nil {
		return err
	}

	return nil
}
