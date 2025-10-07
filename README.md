- Clone the repo: `git clone github.com/not-raghu/go-app`

  **Tip:** If you have just installed [just](https://github.com/casey/just), run:

  - `just dup`
  - `just migrate`
  - `just rundev` (for development)
  - (refer to `justfile` for more commands)

- setup: `docker compose up -d && go run migrate/migrato.go`

- development: `go run main.go` or `air` [anohter compile deamon]
- build: `go build`
