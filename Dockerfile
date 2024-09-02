FROM golang:1.23-alpine as build

WORKDIR /app

COPY . .

RUN go mod download && go test ./...
RUN go build -o /app/autohello ./cmd/autohello/main.go

FROM golang:1.23-alpine

WORKDIR /app

COPY --from=build /app/autohello /usr/local/bin/autohello
