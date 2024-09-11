FROM golang:1.23-alpine as build

WORKDIR /app

COPY . .

RUN go mod download && go test ./...
RUN go build -o /app/bot ./cmd/autohello/main.go

FROM golang:1.23-alpine

WORKDIR /app

COPY --from=build /app/bot /usr/local/bin/bot
