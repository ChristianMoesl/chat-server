FROM golang:1.21.1 as build

WORKDIR /go/src/app
COPY . .

RUN mkdir /app
RUN go mod download \
  && CGO_ENABLED=0 go build -ldflags "-s -w" -o /app/server

RUN mkdir -p /app/database && cp -r ./database/migrations /app/database/migrations \
  && cp -r ./templates /app/templates

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=build /app /app

WORKDIR /app

EXPOSE 8080

CMD ["./server"]
