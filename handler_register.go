package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/rickyjasso/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]
	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}

	_, err := s.db.GetUser(context.Background(), name)
	if err == nil {
		fmt.Println("User already exists")
		os.Exit(1)
	}
	user, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		log.Fatal(err)
	}
	s.cfg.SetUser(user.Name)
	fmt.Printf("User created succesfully. %v\n", user)

	return nil
}
