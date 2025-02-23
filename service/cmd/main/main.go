package main

import (
	"fmt"
	"jrport/minefield/internal/server"
)

func main() {
	server := server.NewGameServer(":8080", 1024)

	if err := server.Run(); err != nil {
		fmt.Println(err.Error())
	}
}
