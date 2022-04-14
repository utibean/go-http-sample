# syntax=docker/dockerfile:1

FROM golang:1.17:alpine
ENV GOPROXY=https://goproxy.cn,direct
ENV HTTP_PORT=8090
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /go-http
EXPOSE 8090
CMD ["/go-http"]
