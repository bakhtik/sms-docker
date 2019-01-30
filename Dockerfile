FROM golang:latest as builder
RUN mkdir -p /app/public
RUN go get -u -v github.com/go-redis/redis
COPY ./*.go /app/
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch
WORKDIR /app
COPY --from=builder /app/main .
COPY public/ /app/public
COPY templates/ /app/templates
EXPOSE 80
CMD ["./main"]