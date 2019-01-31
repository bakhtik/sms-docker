FROM golang:latest as builder
RUN go get -u -v github.com/go-redis/redis
ADD . /app
WORKDIR /app
RUN go build -o main .

# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# FROM scratch
# WORKDIR /app
# COPY --from=builder /app/main .
# COPY public/ /app/public
# COPY templates/ /app/templates
EXPOSE 8080
CMD ["./main"]