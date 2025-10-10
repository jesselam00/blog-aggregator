package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jesselam00/blog-aggregator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err == nil {
		fmt.Fprintf(os.Stderr, "user %s already exists\n", name)
		os.Exit(1)
	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}

	dbUser, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("couldn't create new user: %w", err)
	}

	err = s.cfg.SetUser(dbUser.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("User %s has been created successfully!\n", dbUser.Name)
	log.Printf("Created user: %+v\n", dbUser)
	return nil
}
