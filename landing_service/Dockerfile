FROM golang:1.22-alpine
WORKDIR /app
COPY . ./
COPY *.go ./
COPY templates/* ./templates/
EXPOSE 8080
CMD go run .