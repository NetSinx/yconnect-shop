FROM golang:alpine3.18

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/category .

EXPOSE 8080

CMD ["category"]