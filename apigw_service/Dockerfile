FROM golang:1.22-alpine
WORKDIR /app
COPY . ./
COPY *.go ./
RUN go install github.com/cosmtrek/air@latest
EXPOSE 30200
CMD ["air", "-d","-c", "air.toml"]