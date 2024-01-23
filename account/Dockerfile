FROM golang:1.20-alpine as builder

WORKDIR /go/src/app


RUN go install github.com/cespare/reflex@latest

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./run .


FROM alpine:latest 
WORKDIR /root/
COPY --from=builder /go/src/app/run .

EXPOSE 8080
CMD ["./run"]