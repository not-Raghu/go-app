dup:
    @echo "starting docker containers"
    @docker compose up -d
dstop:
    @echo "stopping containers"
    @docker compose stop
ddown:
    @echo "bringing down containers"
    @docker compose down 
dflush:
    @echo "bringing down containers and flushing the data"
    @docker compose down -v
migrate:
    @echo "running db migration"
    @go run migrate/migrate.go
seed:
    @echo "seeding database"
    @go run migrate/migrate.go seed
run:
    @echo "running the app"
    @bash createfile.sh
    @go run main.go
rundev:
    @echo "running app in dev"
    @bash createfile.sh
    @air
# build:
#     @echo "building app"
#     @go build -o ./dist 
# ga:
#     @echo "git add ."

#add tests here later
