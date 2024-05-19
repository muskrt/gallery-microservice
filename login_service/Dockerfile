FROM golang:1.22-alpine
WORKDIR /app
COPY . ./
COPY *.go ./
RUN go install github.com/cosmtrek/air@latest
EXPOSE 30100
CMD ["air", "-d","-c", "air.toml"]