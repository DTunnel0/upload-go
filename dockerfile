FROM golang:1.21

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o app src/main.go  

CMD ["./app"]
