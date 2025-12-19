package main

import (
	"context"
	"fmt"

	"github.com/rickyjasso/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CURRENT_USER_NAME)
		if err != nil {
			return fmt.Errorf("Error getting user: %w", err)
		}
		handler(s, cmd, user)
		return nil
	}
}
