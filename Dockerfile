FROM golang

WORKDIR /go/src/app

COPY . .

RUN go mod download
RUN go mod verify
RUN go build -o todo-api

CMD ["./todo-api"]