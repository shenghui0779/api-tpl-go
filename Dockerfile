FROM golang:1.16.5 AS builder

WORKDIR /tplgo

COPY . .

RUN go env -w GOPROXY="https://goproxy.cn"

RUN go mod tidy

RUN sh ent.sh

RUN CGO_ENABLED=0 go build -o ./bin/main ./cmd

FROM scratch

WORKDIR /data

COPY --from=builder /tplgo/bin/main .

EXPOSE 8000

ENTRYPOINT ["./main"]

CMD ["--envfile", "/data/config/.env"]