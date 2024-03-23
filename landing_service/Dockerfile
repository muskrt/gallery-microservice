FROM golang:1.22-alpine
WORKDIR /landing_service
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go build -o .
EXPOSE 8080
CMD [ "landing_service" ]