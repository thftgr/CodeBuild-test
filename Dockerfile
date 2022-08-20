FROM golang:alpine AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY go.mod go.sum main.go ./
RUN go mod download
RUN go build -o main .
WORKDIR /dist
RUN cp /build/main .

FROM scratch
COPY --from=builder /dist/main .

ENV MYSQL_USER=user
ENV MYSQL_PASSWORD=password
ENV MYSQL_HOST=mysql
ENV MYSQL_PORT=3306
ENV MYSQL_SCHEMA=""
ENV SERVER_PORT=80

EXPOSE 80
ENTRYPOINT ["/main"]