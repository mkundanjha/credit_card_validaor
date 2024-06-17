# FROM arm64v8/golang:latest
FROM arm64v8/golang:1.21rc4-alpine

WORKDIR /app/src

COPY . /app/src

RUN go mod tidy

RUN go build -o ccvalidator .

ENTRYPOINT ["./ccvalidator"]

