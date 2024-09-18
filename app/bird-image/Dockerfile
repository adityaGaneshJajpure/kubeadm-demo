FROM golang:alpine as builder

RUN apk update && apk add --no-cache git
RUN apk add --update --no-cache openssh

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download 
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/bird-image

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /app/main .
RUN chown -R appuser:appgroup /app

USER appuser

EXPOSE 8000

CMD ["./main"]
