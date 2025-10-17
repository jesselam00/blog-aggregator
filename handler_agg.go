package main

import (
	"context"
	"fmt"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jesselam00/blog-aggregator/internal/database"
)

func handlerAgg(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 1 || len(cmd.Args) > 2 {
		return fmt.Errorf("usage: %v <time_between_reqs>", cmd.Name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("incorrect duration string: %s <time_between_requests>", cmd.Name)
	}

	ticker := time.NewTicker(timeBetweenRequests)

	fmt.Printf("Collection feeds every %v...\n", timeBetweenRequests)

	for ; ; <-ticker.C {
		err = scrapeFeeds(s, user.ID)
		if err != nil {
			return fmt.Errorf("couldn't fetch feed: %w", err)
		}
	}
}

func scrapeFeeds(s *state, userId uuid.UUID) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background(), userId)
	if err != nil {
		return fmt.Errorf("could not find next feed to fetch: %w", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID: feed.ID,
		LastFetchedAt: sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		}, 
	})
	if err != nil {
		return fmt.Errorf("could not mark feed as fetched: %w", err)
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("could not fetch feed: %w", err)
	}

	fmt.Println("============================")
	fmt.Println(rssFeed.Channel.Title)
	for _, rssItem := range rssFeed.Channel.Item {
		fmt.Printf(" - %s\n", rssItem.Title)
	}
	fmt.Println("============================")

	return nil
}