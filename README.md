# 🐊 Blog Aggregator

**Blog Aggregator** is a microservice CLI RSS feed aggregator written in Go! We'll call the CLI tool `gator` because it is an aggreGATOR 🐊. This app allows you to:

* Add RSS feeds from across the internet.
* Store collected posts in a PostgreSQL database.
* Follow and unfollow RSS feeds from other users.
* View summaries of aggregated posts in the terminal, with links to full posts.

RSS feeds are a convenient way to keep up with blogs, news sites, podcasts, and more.

---

## Prerequisites

To run this project, you need:

* [Go 1.25+](https://go.dev/doc/install)
* [PostgreSQL 15+](https://www.postgresql.org/download/)

---

## Installing the CLI

You can install the CLI using:

```bash
go install github.com/yourusername/blog-aggregator@latest
```

This will install a `blog-aggregator` executable you can use in your terminal.

---

## Configuration

`gator` uses a JSON config file located at `~/.gatorconfig.json`:

```json
{
  "db_url": "postgres://user:password@localhost:5432/blog_aggregator",
  "current_user_name": "your_username"
}
```

* `db_url`: Connection string to your PostgreSQL database.
* `current_user_name`: Automatically set by the application after using the register.

---

## Running the CLI

For development, you can run:

```bash
go run .
```

For production usage, use the installed `blog-aggregator` binary:

```bash
blog-aggregator <command> [args]
```

---

## Commands

The Blog Aggregator CLI supports the following commands:

| Command    | Description                                         | Usage                          | Login Required |
|------------|-----------------------------------------------------|--------------------------------|----------------|
| register   | Create a new user                                   | `register <name>`              | No             |
| login      | Switch to an existing user                          | `login <name>`                 | No             |
| reset      | Reset the database by deleting all users            | `reset`                        | No             |
| users      | List all users and highlight the current one       | `users`                        | No             |
| agg        | Collect feeds periodically                          | `agg <time_between_requests>`  | Yes            |
| addfeed    | Add a new RSS feed                                   | `addfeed <feed_url>`           | Yes            |
| feeds      | List all available feeds                             | `feeds`                        | No             |
| follow     | Follow an existing feed                              | `follow <feed_url>`            | Yes            |
| following  | List all feeds the current user follows             | `following`                    | Yes            |
| browse     | Browse posts for the current user (optional limit)  | `browse [limit]`               | Yes            |


---

## Database Setup

PostgreSQL is used to store users, feeds, and posts.

### macOS (Homebrew)

```bash
brew install postgresql@15
brew services start postgresql@15
psql postgres
```

### Linux / WSL (Debian)

```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
sudo service postgresql start
sudo -u postgres psql
```

### Create Database

```sql
CREATE DATABASE gator;
\c gator
ALTER USER postgres PASSWORD 'postgres';
```

---
