FROM golang:1.22-alpine
WORKDIR /app
RUN rm -rf ./tmp
COPY . ./
COPY *.go ./
RUN go install github.com/pilu/fresh@latest
COPY templates/* ./templates/
EXPOSE 8080
CMD fresh