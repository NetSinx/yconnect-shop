FROM golang:alpine

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/user ./cmd/

EXPOSE 8082

CMD ["user"]
