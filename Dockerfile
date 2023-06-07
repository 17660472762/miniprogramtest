FROM golang:1.18 as builder

RUN mkdir /app

ADD . /app/

WORKDIR /app


RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

RUN echo 'Asia/Shanghai' >/etc/timezone

RUN go env -w GOPROXY=https://goproxy.cn

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/mini_program/main.go

FROM alpine:latest

RUN apk add ca-certificates

WORKDIR /app

COPY --from=builder /app/main .

CMD ["/app/main","-appid","wxde659347f55b799b","-secret","9a022075538836978d7b3003b72bfcfd"]
