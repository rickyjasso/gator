package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Error getting feeds: %w\n", err)
	}

	fmt.Println("All Feeds:")
	fmt.Println("===============")
	for _, feed := range feeds {
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("Error getting user: %w\n", err)
		}
		fmt.Printf("Feed: %s\n", feed.Name)
		fmt.Printf("- URL: %s\n", feed.Url)
		fmt.Printf("- Creator: %s\n", user.Name)
		fmt.Println()
	}
	return nil

}
