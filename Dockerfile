FROM golang:1.21.4 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./server ./src/main.go 

FROM ubuntu:focal
WORKDIR /app
COPY --from=builder /app/server .

RUN apt-get update -y
RUN apt-get upgrade -y

RUN apt-get install git -y
RUN apt-get install -y openssh-client
RUN apt-get install sshpass -y
RUN apt-get install curl -y

EXPOSE 80

CMD ["/app/server"]

