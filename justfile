#(fuck makefile)

dup:
    @echo "starting docker containers"
    @docker compose up -d
ddown:
    @echo "bringing down containers"
    @docker compose down 
dflush:
    @echo "bringing down containers and flushing the data"
    @docker compose down -v
migrate:
    @echo "running db migration"
    @go run migrate/*.go
seed:
    @echo "seeding database"
    @go run seed.go
run:
    @echo "running the app"
    @go run main.go
rundev:
    @echo "running app in dev"
    @air

#add tests here later