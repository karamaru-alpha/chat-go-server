FROM golang:1.16

WORKDIR /go/src/github.com/karamaru-alpha/chat-go-server

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build .

RUN go get github.com/pilu/fresh

EXPOSE 8080

CMD ["fresh"]
