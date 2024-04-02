FROM golang:1.22-alpine
WORKDIR /app
COPY . ./
COPY *.go ./
COPY *.go ./
RUN go install github.com/pilu/fresh@latest
EXPOSE 30100
CMD fresh