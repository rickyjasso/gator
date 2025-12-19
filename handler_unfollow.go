package main

import (
	"context"
	"fmt"

	"github.com/rickyjasso/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: %s <url>\n", cmd.Name)
	}

	url := cmd.Args[0]

	feedToUnfollow, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Error getting feed to unfollow: %w", err)
	}

	params := database.UnfollowFeedParams{
		UserID: user.ID,
		FeedID: feedToUnfollow.ID,
	}

	_, err = s.db.UnfollowFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("Error unfollowing feed: %w", err)
	}

	fmt.Printf("Succesfully unfollowed feed: %s\n", feedToUnfollow.Name)
	return nil
}
