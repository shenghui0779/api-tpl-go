FROM golang:1.16.4 AS builder

WORKDIR /tplgo

COPY . .

RUN CGO_ENABLED=0 go build -mod=vendor -o tplgo ./cmd

FROM scratch

WORKDIR /goapp

COPY --from=builder /tplgo/tplgo .

EXPOSE 10086

ENTRYPOINT ["./tplgo"]