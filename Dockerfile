FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./src ./src
RUN go build ./src/cmd


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
CMD ["./cmd"]  