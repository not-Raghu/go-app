# WHAT IS THIS? ID

prolly a blog site like medium to learn go, idk, drop some suggestions if you have pls

## reproducing project locally

Tip: If you have [just](https://github.com/casey/just) installed, run:

- USING JUST :

  - `just dup`
  - `just migrate`
  - `just rundev` (for development)
    <br>
    (refer to just file for more commands)
    </br>

- METHOD 2

  - `bash createfile.sh`
  - env: `add your environment variables`
  - setup: `docker compose up -d && go run migrate/migrato.go`

  - development: `go run main.go` or `air` [or any other compile deamon]
  - build: `go build`
