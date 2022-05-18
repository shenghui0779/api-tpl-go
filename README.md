# tplgo

A template for building go web service with [yiigo](https://github.com/shenghui0779/yiigo).

> 1. ORM [entgo.io](https://entgo.io/)
> 2. Table `User` refers to `pkg/ent/schema/user.go`
> 3. Set `GOPROXY` [ `go env -w GOPROXY="https://goproxy.cn,direct"` ]

### 1. prepare

```shell
go mod download
sh ent.sh
go mod tidy -compat=1.17
```

### 2. run

```shell
mv .env.example => cmd/.env
cd cmd          => go run main.go
```
