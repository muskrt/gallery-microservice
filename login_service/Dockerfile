FROM golang:1.22-alpine
WORKDIR /login_service
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go build -o .
EXPOSE 30300
CMD [ "login_service" ]