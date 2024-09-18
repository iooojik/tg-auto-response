FROM golang:1.23-alpine as build

COPY . .

RUN go mod download && go test ./...
RUN go build -o /app/bot /app/cmd/autohello/main.go

FROM golang:1.23-alpine

COPY --from=build /app/bot /usr/local/bin/bot
