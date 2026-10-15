FROM golang:1.27

WORKDIR /go/src/app

COPY . . 

EXPOSE 8000

# Build the Go app
RUN go build -o main cmd/main.go

CMD ["./main"]