package main

import (
	"context"
	"fmt"

	"github.com/rickyjasso/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {

	following, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Error getting the user's followed feeds: %w", err)
	}

	printFollowingFeeds(following)

	return nil
}

func printFollowingFeeds(feeds []database.GetFeedFollowsForUserRow) {
	fmt.Println("Followed feeds")
	fmt.Println("==============")
	for _, feed := range feeds {
		fmt.Printf("* Feed: %v\n", feed.FeedName)
		fmt.Printf("* Created At: %v\n", feed.CreatedAt)
		fmt.Printf("* User: %v\n", feed.UserName)
		fmt.Println()
	}
}
