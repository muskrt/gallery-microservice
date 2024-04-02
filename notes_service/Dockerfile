FROM golang:1.22-alpine
WORKDIR /app
COPY . ./
COPY *.go ./
RUN go install github.com/pilu/fresh@latest
EXPOSE 30300
CMD fresh