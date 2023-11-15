FROM golang:1.21

WORKDIR /app

COPY project ./project

WORKDIR /app/project

RUN go mod download

EXPOSE 80

ENTRYPOINT go run .
