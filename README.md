# Chat Server

Simple chat server, where people can send messages

## Setup
* Install Planetscale CLI tool for development
```
brew install planetscale/tap/pscale
brew install mysql-client
```

* Install migration tool to generate SQL migrations
```
go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate
```

* Install AIR for hot reloading:
`go install github.com/cosmtrek/air@latest`

* Run server in development mode
`MODE=development air`

## Deploy

* Deploy to Google Cloud Run:
`gcloud run deploy chat-server --source . --allow-unauthenticated --region=europe-west3`
