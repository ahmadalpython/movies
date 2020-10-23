FROM golang:1.15.3-alpine3.12

WORKDIR /go/src/app
COPY . .
RUN go build -o main .

CMD ["./main"]