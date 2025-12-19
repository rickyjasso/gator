# Gator 🐊

Gator is a command-line RSS feed aggregator written in Go. It uses PostgreSQL for data storage and allows users to register, follow feeds, and browse aggregated content from the terminal.

## Requirements

Before running Gator, you must have the following installed:

- Go (1.20 or newer)
  https://go.dev/doc/install

- PostgreSQL
  https://www.postgresql.org/download/

A running PostgreSQL database is required.

## Installation

Install the Gator CLI using Go:

>go install github.com/rickyjasso/gator@latest

Ensure your Go bin directory is in your PATH:

export PATH=$PATH:$(go env GOPATH)/bin

After this, the `gator` command should be available.

## Configuration

Gator uses a config file stored in your home directory.

### Create the config file

touch ~/.gatorconfig.json

### Example configuration

{\
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",\
  "current_user": ""\
}

- db_url is your PostgreSQL connection string
- current_user is automatically set when you log in

Make sure the database exists before running the program.

## Running the Program

Run commands using the following format:

gator <command> [arguments]
