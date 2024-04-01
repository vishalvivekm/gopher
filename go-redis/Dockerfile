FROM golang

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go get get -t
RUN go build -o main .

EXPOSE 9090
CMD ["/app/main"]