# Chat Server

Simple chat server, where people can send messages

## Setup
```
* Use AIR for hot reloading:
`go install github.com/cosmtrek/air@latest`

* Run server in development mode
`MODE=development air`

## Deploy

* Deploy to Google Cloud Run:
`gcloud run deploy chat-server --source . --allow-unauthenticated --region=europe-west1`
