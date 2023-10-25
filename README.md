# api-tpl-go

Go API 项目框架 ( [yiigo](https://github.com/shenghui0779/yiigo) + [ent](https://entgo.io) + [chi](https://github.com/go-chi/chi) ) 👉 你想要的基本都有

> 1. Table `User` refers to `ent/schema/user.go`
> 2. Set `GOPROXY` ( `go env -w GOPROXY="https://proxy.golang.com.cn,direct"` )

### 1. prepare

```shell
go mod download
sh ent.sh
go mod tidy
```

### 2. run

```shell
mv .env.example env
go run main.go
```
