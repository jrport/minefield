package main

import "fmt"

func MatchMakingService(um *UserManager) {
	for {
		_ = <-um.Notifier
		if um.UserCount < 2 {
			continue
		}

		p1 := um.Dequeue()
		p2 := um.Dequeue()
		fmt.Printf(`
			Game ready for players:\n
			Player 1: {
				name: %v,
				id: %v,
			},
			Player 2: {
				name: %v,
				id: %v,
			},
			`, p1.Name, p1.Id, p2.Name, p2.Id,
		)
	}
}
