FROM golang:1.17.8-buster
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN go build -o main .
CMD ["/app/main"]
