# 🐊 Blog Aggregator

Welcome to **Blog Aggregator** — a microservice command-line RSS feed aggregator written in Go. It collects and stores posts from multiple RSS feeds in a PostgreSQL database, allowing you to fetch, follow, and read posts all in one place.

---

## Prerequisites

Before you get started, make sure you have:

* **Go** installed (version 1.22+ recommended)
  [Download Go](https://go.dev/dl/)

* **PostgreSQL** installed and running
  [Download PostgreSQL](https://www.postgresql.org/download/)

You’ll also need a Postgres database set up for this project. Example:

```bash
createdb gator
```

---

## Installation

To install the CLI tool (`gator`), run:

```bash
go install github.com/jesselam00/blog-aggregator@latest
```

Make sure your Go bin is on your PATH:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

Now you can run the CLI anywhere by typing:

```bash
gator
```

---

## Configuration

`gator` requires a config file to connect to your database and manage user sessions.

Run the following command to set up your config file:

```bash
gator register <username>
```

This creates a config file in:

```
~/.gatorconfig.json
```

Inside it, you’ll find your database connection string and the current user.

---

## Usage

Here are a few useful commands:

| Command                      | Description                                               |
| ---------------------------- | --------------------------------------------------------- |
| `gator register <username>`  | Registers a new user and creates the config file          |
| `gator login <username>`     | Logs in as an existing user                               |
| `gator addfeed <name> <url>` | Adds a new RSS feed to track                              |
| `gator listfeeds`            | Lists all feeds currently being tracked                   |
| `gator agg`                  | Starts fetching posts from all feeds in a continuous loop |
| `gator browse`               | Displays the most recent posts for your user              |

You can always run:

```bash
gator help
```

to see all available commands.

---

## Development

To build locally (for testing or development):

```bash
go run .
```

For production use (after building/installing):

```bash
gator
```

Go programs are statically compiled — once built, you can run the `gator` binary without needing Go installed on the target machine.

---

## Database Setup

Before running `gator`, ensure your PostgreSQL database is ready:

1. Start your Postgres server.
2. Run the migrations:

   ```bash
   goose -dir sql/schema postgres "postgres://<user>:<password>@localhost:5432/gator" up
   ```
3. Confirm tables were created:

   ```bash
   psql gator
   \dt
   ```

---

## Repository

GitHub Repository:
[https://github.com/jesselam00/blog-aggregator](https://github.com/jesselam00/blog-aggregator)

---

Happy hacking! 🐊
Now go aggregate some blogs!
