# api-tpl-go

Go API 快速开发脚手架

> 1. `config.toml.example` => `config.toml`
> 2. Table `User` refers to `ent/schema/user.go`
> 3. 执行 `ent.sh` 生成ORM代码 (只要 `ent/schema` 目录下有变动都需要执行)
> 4. Set `GOPROXY` ( `go env -w GOPROXY="https://proxy.golang.com.cn,direct"` )

- 路由使用 [chi](https://github.com/go-chi/chi)
- ORM使用 [ent](https://github.com/ent/ent)
- Redis使用 [go-redis](https://github.com/redis/go-redis)
- 日志使用 [zap](https://github.com/uber-go/zap)
- 配置使用 [viper](https://github.com/spf13/viper)
- 命令行使用 [cobra](https://github.com/spf13/cobra)
- 工具包使用 [yiigo](https://github.com/shenghui0779/yiigo)
- 包含 认证、请求日志、跨域 中间价
- 简单好用的 API Result 统一输出方式
