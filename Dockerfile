RUN useradd -u 8877 user_runner
USER user_runner

FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./src ./src
RUN go build ./src/main.go


FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app ./
ENV HOST=postgres
ENV USERNAME=postgres
ENV PASSWORD=postgres
ENV DATABSE=gorm
ENV DB_PORT=5432
ENV SDSLMODE=disable
ENV TIMEZONE=UTC
CMD ["./main"]  