package main

import (
	"context"
	"fmt"
	"log"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		if s.cfg.CURRENT_USER_NAME == user.Name {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}
	return nil
}
