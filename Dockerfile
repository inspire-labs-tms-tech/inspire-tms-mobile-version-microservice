FROM --platform=linux/amd64 golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server .

FROM --platform=linux/amd64 alpine:3.21

WORKDIR /app
RUN apk --no-cache add ca-certificates

# copy the compiled binary from the builder stage
COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]