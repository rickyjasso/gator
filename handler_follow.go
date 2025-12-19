package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rickyjasso/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: %s <url>\n", cmd.Name)
	}

	url := cmd.Args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Error getting feed by URL: %w", err)
	}

	feedFollowsParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), feedFollowsParams)
	if err != nil {
		return fmt.Errorf("Error creating new feed follow: %w", err)
	}

	fmt.Printf("Created new feed follow\n")
	fmt.Printf("- User: %s\n", user.Name)
	fmt.Printf("- Feed: %s\n", feed.Name)

	return nil

}
