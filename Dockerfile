FROM golang:1.17.8 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o rest-go

FROM scratch
COPY --from:builder /app/rest-go .
ENTRYPOINT ["./rest-go"]
