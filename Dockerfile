FROM golang:latest

WORKDIR /go/src/github.com/andreoav/shopbot

COPY . .

RUN go get github.com/pilu/fresh

RUN go-wrapper download

EXPOSE 8080

CMD ["fresh"]
