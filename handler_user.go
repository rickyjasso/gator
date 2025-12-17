package main

import (
	"context"
	"fmt"
	"os"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	err := s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set user: %w", err)
	}

	_, err = s.db.GetUser(context.Background(), name)
	if err != nil {
		fmt.Printf("User does not exist.")
		os.Exit(1)
	}

	fmt.Println("User switched successfully")
	return nil
}
