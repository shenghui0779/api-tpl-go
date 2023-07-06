# tplgo

Go web api 开发模板 ( [yiigo](https://github.com/shenghui0779/yiigo) + [chi](https://github.com/go-chi/chi) ) 👉 你想要的基本都有

> 1. ORM [entgo.io](https://entgo.io/)
> 2. Table `User` refers to `pkg/ent/schema/user.go`
> 3. Set `GOPROXY` ( `go env -w GOPROXY="https://goproxy.cn,direct"` )

### 1. prepare

```shell
go mod download
sh ent.sh
go mod tidy
```

### 2. run

```shell
mv .env.example => cmd/.env
cd cmd          => go run main.go
```
