# chirpy

## What is it?

**chirpy** is an implementation of the guided project [Build a Social Media API](https://www.boot.dev/courses/build-social-media-api) from the [boot.dev](https://www.boot.dev) platform.

It is a social media API that allows users to post messages (chirps), and view chirps posted by other users.

## Requirements

- Go, documentation to install Go can be found [here](https://go.dev/doc/install)
- PostgreSQL, documentation to install PostgreSQL can be found [here](https://www.postgresql.org/download/)

## Installation

The source code is currently hosted on [GitHub](https://www.github.com/noch-g/chirpy).

```bash
> go install github.com/noch-g/chirpy@latest
```
Make sure you create a `chirpy` database on your PostgreSQL server:
```bash
> createdb chirpy
```

## Configuration
The application needs to be configured with variables place in a `.env` file.
The `.env` expects the following variables:

```bash
# File: .env
DB_URL="postgres://<user>:@localhost:5432/chirpy?sslmode=disable"
TOKEN_SECRET="your-secret-token-here"
POLKA_KEY="your-polka-key-here"
PLATFORM="dev"
```
- In `DB_URL`, replace `<user>` with the user you want to use to connect to the database, and `chirpy` with the name of the database if you chose to name it differently.
- `TOKEN_SECRET` should be a random string of your choice, it is used to sign and verify JWT tokens.
- `POLKA_KEY` is the API key for the (fake) Polka payment processor, it is used to authenticate calls made to the webhook endpoint.
- `PLATFORM` should be set to `dev` in order to run the application locally, and `prod` when deploying to a production environment.
