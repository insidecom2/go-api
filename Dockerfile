FROM golang:1.17-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go mod download github.com/ugorji/go
RUN go build -o ./go-api ./server.go


FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/go-api .
EXPOSE 8000
ENTRYPOINT ["./go-api"]