FROM golang:1.22-alpine
WORKDIR /app
COPY . ./
COPY *.go ./
RUN go install github.com/pilu/fresh@latest
COPY templates/* ./templates/
EXPOSE 8080
CMD fresh