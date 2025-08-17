FROM golang:1.25.0-alpine

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/auth .

EXPOSE 8086

CMD ["product"]