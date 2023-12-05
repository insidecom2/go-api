FROM golang:1.17-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o go-api

FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/go-api /app/go-api
COPY .env.production ./.env
ENV $(cat .env | xargs)
EXPOSE 8080

CMD [ "./go-api" ]