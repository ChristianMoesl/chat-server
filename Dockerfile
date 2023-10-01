FROM golang:1.21.1 as build

WORKDIR /go/src/app
COPY . .

RUN go mod download \
  && CGO_ENABLED=0 go build -ldflags "-s -w" -o /go/bin/app cmd/server/server.go

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=build /go/src/app/database/migrations /database/migrations
COPY --from=build /go/bin/app /

EXPOSE 8080

CMD ["/app"]
