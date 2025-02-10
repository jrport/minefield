package server

import (
	"fmt"
	"jrport/minefield/internal/pool"
	"net/http"
	//"strings"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

type MatchMakerHandle struct {
	matchPool   *pool.MatchPool
	redisClient *redis.Client
}

func (mh *MatchMakerHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("im here")

	//token := strings.TrimSpace(r.Header.Get("Authorization"))

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	// TODO use this name in matchmaking
	//_, err := mh.redisClient.Get(ctx, token).Result()
	//if err == redis.Nil {
	//	http.Error(w, "Invalid Token", http.StatusUnauthorized)
	//	return
	//} 
	//if err != nil {
	//	http.Error(w, "Internal error", http.StatusInternalServerError)
	//	return
	//} 

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// TODO
	// Check for wether the matchpool has matched this guy to someone
	// wait for that looping the standby msg 
	// signal when ready to go to the next step

	for {
		messageType, _, err := conn.NextReader()
		if err != nil {
			return
		}
		w, err := conn.NextWriter(messageType)
		if err != nil {
			return
		}
		fmt.Fprint(w, "HOLD")
		w.Write([]byte("HODL"))
		if err := w.Close(); err != nil {
			return
		}
	}
}

func newMatchMakerHandler(mp *pool.MatchPool, rc *redis.Client) *MatchMakerHandle {
	return &MatchMakerHandle{
		matchPool:   mp,
		redisClient: rc,
	}
}
