FROM golang:1.8

EXPOSE 8080
WORKDIR /go/src/app
COPY . .

RUN go build -o go-api main.go

CMD ["/go/src/app/go-api"]

