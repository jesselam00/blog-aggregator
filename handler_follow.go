package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jesselam00/blog-aggregator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't follow feed: %w", err)
	}

	fmt.Printf("%s is now following feed %s\n", user.Name, feed.Name)
	return nil
}

func handlerFollowing(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get followed feeds: %w", err)
	}

	fmt.Printf("%s is following %d feeds:\n", user.Name, len(feedFollows))
	for _, ff := range feedFollows {
		feed, err := s.db.GetFeedById(context.Background(), ff.FeedID)
		if err != nil {
			return fmt.Errorf("couldn't get feed: %w", err)
		}
		fmt.Printf("- %s\n", feed.Name)
	}
	return nil
}
