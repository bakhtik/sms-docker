FROM golang:latest as builder
RUN go get -u -v github.com/go-redis/redis
ADD . /go/src/github.com/bakhtik/sms-docker
WORKDIR /go/src/github.com/bakhtik/sms-docker
RUN go build -o main cmd/sms-docker/main.go

# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# FROM scratch
# WORKDIR /app
# COPY --from=builder /app/main .
# COPY public/ /app/public
# COPY templates/ /app/templates
EXPOSE 8080
CMD ["./main"]