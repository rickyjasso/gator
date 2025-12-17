package main

import (
	"github.com/rickyjasso/gator/internal/config"
	"github.com/rickyjasso/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}
