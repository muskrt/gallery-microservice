FROM golang:1.22-alpine
WORKDIR /app
COPY . ./
COPY *.go ./
COPY templates/* ./templates/
RUN go install github.com/cosmtrek/air@latest
EXPOSE 8080
CMD ["air", "-d","-c", "air.toml"]