FROM golang:1.22-alpine
WORKDIR /login_service
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY *.go ./
EXPOSE 30300
CMD go run .