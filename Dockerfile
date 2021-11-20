FROM golang:1.16.5 AS builder

WORKDIR /tplgo

COPY . .

RUN go env -w GOPROXY="https://goproxy.cn"

RUN go mod tidy

RUN CGO_ENABLED=0 go build -o ./bin/main ./cmd

FROM scratch

WORKDIR /data

COPY --from=builder /tplgo/bin/main ./bin

EXPOSE 10086

ENTRYPOINT ["./bin/main"]

CMD ["--envfile", "/data/config/.env"]