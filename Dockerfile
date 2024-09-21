FROM golang:1.23-alpine AS build

COPY . .

RUN go mod download && go test ./...
RUN go build -o /app/bot cmd/autoresponse/main.go

FROM golang:1.23-alpine

COPY --from=build /app/bot /usr/local/bin/bot
