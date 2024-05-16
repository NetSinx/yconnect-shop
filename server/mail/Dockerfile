FROM golang:alpine3.18

WORKDIR /usr/src/app

COPY . .

RUN go mod download && go mod verify && go build -v -o /usr/local/bin/mail .

EXPOSE 8085

CMD ["mail"]