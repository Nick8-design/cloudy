# Use the official Golang image as a build stage
FROM golang:1.23.5


WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY main.go .

RUN go build -o bin .


ENTRYPOINT ["/app/bin"]
