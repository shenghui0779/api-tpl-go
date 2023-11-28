FROM golang:1.19.4 AS builder

WORKDIR /api

COPY . .

RUN go env -w GOPROXY="https://proxy.golang.com.cn,direct"
RUN go mod download
RUN sh ent.sh
RUN go mod tidy

RUN CGO_ENABLED=0 go build -o ./bin/main

FROM scratch

WORKDIR /bin

COPY --from=builder /api/bin/main .

EXPOSE 8000

ENTRYPOINT ["./main"]

CMD ["--config", "/data/config/.yml"]
