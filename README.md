## Libraries used
- Configuration: https://github.com/spf13/viper
- HTTP Router: https://github.com/gorilla/mux

## Run

1. Start database via Docker

`docker run --name buildit-db -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 postgres`
   
2. Start service 
   
`go run main.go http -l debug`
or via Docker see [Docker](#docker)

Database migrations will be performed at first execution.

## Run tests

`go test ./test/e2e`

## Docker

Build
`docker build -t buildit:latest --build-arg SSH_PRIVATE_KEY="$(cat ~/.ssh/id_rsa)" .`

Run
`docker run -p 8080:8080 --network host buildit:latest`