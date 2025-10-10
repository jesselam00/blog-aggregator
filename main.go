package main

import (
	"fmt"

	"github.com/jesselam00/blog-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Error reading config:", err)
		return
	}

	cfg.SetUser("jesse")

	cfg, err = config.Read()
	if err != nil {
		fmt.Println("Error reading config:", err)
		return
	}

	fmt.Println("Configuration:")
	fmt.Println("db_url:", cfg.DbUrl)
	fmt.Println("current_user_name:", cfg.CurrentUser)

}
