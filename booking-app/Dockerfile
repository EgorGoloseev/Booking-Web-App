FROM golang:1.22-alpine
WORKDIR /app
COPY go.mod go.sum ./
COPY . .
RUN go build -mod=vendor -o server .
EXPOSE 8080
CMD ["./server"]