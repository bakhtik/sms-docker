FROM golang:latest as builder
RUN mkdir -p /app/public
COPY ./*.go /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch
WORKDIR /app
COPY --from=builder /app/main .
COPY public/ /app/public
COPY templates/ /app/templates
CMD ["./main"]