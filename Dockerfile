FROM golang:1.16.5 AS builder

WORKDIR /tplgo

COPY . .

RUN go env -w GOPROXY="https://goproxy.cn"

RUN go mod tidy

RUN CGO_ENABLED=0 go build -o tplgo ./cmd

FROM scratch

WORKDIR /goapp

COPY --from=builder /tplgo/tplgo .

EXPOSE 10086

ENTRYPOINT ["./tplgo"]

CMD ["--env-dir", "/data/config"]