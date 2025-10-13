# jwt-auth-gin

Simple starter project showing a Gin HTTP server with JWT auth initializers.

## What this repo contains

- `main.go` — application entry. Calls `initializers.LoadEnvVariable()` from `init()` and starts a Gin server with a `/ping` route.
- `initializers/` — package for loading environment variables and DB connection (existing files: `loadEnvVariable.go`, `connectToDb.go`).

## Run locally

Quick run (no build):

```bash
# run the app directly
go run main.go
```

Build and run:

```bash
# build binary
go build -o jwt-auth-gin
# run binary
./jwt-auth-gin
```

Optional: use a live-reload tool during development. Example with CompileDaemon:

```bash
# install (adds binary to GOPATH/bin)
go install github.com/githubnemo/CompileDaemon@latest
# make sure GOPATH/bin is on your PATH, e.g. add to ~/.zshrc
export PATH="$PATH:$(go env GOPATH)/bin"
# then run
CompileDaemon --command="./jwt-auth-gin"
```

Alternatives: `air`, `reflex` (install via `go install github.com/cosmtrek/air@latest` or `go install github.com/cespare/reflex@latest`).

## Notes

- `init()` in `main.go` calls `initializers.LoadEnvVariable()` — ensure any required `.env` exists or update the initializer.
- The repo currently exposes a test `/ping` endpoint that returns `{"message":"pong"}`.

## Next steps

- Add `.env.example` with required variables.
- Add a `Makefile` or `task` for common commands.
- Add basic unit/integration tests for the initializers and handlers.
