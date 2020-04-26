FROM golang:1.13

RUN mkdir /app

ADD . /app/

WORKDIR /app

RUN GO111MODULE=on go build .

CMD ["/app/warung_pintar"]