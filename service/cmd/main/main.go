package main

import (
	"fmt"
	"jrport/minefield/internal/server"
)

func main() {
	server := server.NewGameServer(":3000", 1024)

	if err := server.Run(); err != nil {
		fmt.Println(err.Error())
	}
}
