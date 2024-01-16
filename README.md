# Golang: Basic Gin Web Framework Example

### Example for students in 2110336 SOFTWARE ENGINEERING II Class

### Disclaimer: This is a basic example for beginners and may not include technique or method such as

- Standard Go Project Layout
- Advance Env Management
- Complex Design Pattern or Architecture
- Hot Reload
- Etc.

Note: This repository is one of many possible ways to create a web server with gin in golang, other project structures or methods are fine too.

## Requirement

- Golang
- Docker with Docker Compose

## Usage

- copy `.env.template` to `.env` and configure the environment variable.
- run `docker compose up` to start up postgres database.
- run `go run ./main.go` to start up an application.

## Q&A

- Why Golang?

Golang is a compiler language, so performance is better than interpreter language in general such as Python or Javascript, Golang also has built-in concurrency support through Goroutines.

- Why Gin?

Gin's community is larger so it is easier for beginners to find reference or information.
