package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/rickyjasso/gator/internal/config"
	"github.com/rickyjasso/gator/internal/database"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	s := state{
		cfg: &cfg,
	}

	db, err := sql.Open("postgres", s.cfg.DB_URL)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)
	s.db = dbQueries

	commands := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)
	commands.register("agg", handlerAggregator)
	commands.register("addfeed", middlewareLoggedIn(handlerFeed))
	commands.register("feeds", handlerFeeds)
	commands.register("follow", middlewareLoggedIn(handlerFollow))
	commands.register("following", middlewareLoggedIn(handlerFollowing))
	commands.register("unfollow", middlewareLoggedIn(handlerUnfollow))

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Error! Need arguments")
		os.Exit(1)
	}

	cmd := command{
		Name: args[1],
		Args: args[2:],
	}

	err = commands.run(&s, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
