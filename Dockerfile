FROM golang:1.17.10 AS builder

WORKDIR /tplgo

COPY . .

RUN go env -w GOPROXY="https://goproxy.cn,direct"
RUN go mod download
RUN sh ent.sh
RUN go mod tidy -compat=1.17

RUN CGO_ENABLED=0 go build -o ./bin/main ./cmd

FROM scratch

WORKDIR /bin

COPY --from=builder /tplgo/bin/main .

EXPOSE 8000

ENTRYPOINT ["./main"]

CMD ["-envfile", "/data/config/.env"]
