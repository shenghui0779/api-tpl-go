package cmd

import (
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

func Init() {
	// 注册命令
	root.AddCommand(helloCmd)

	// 注册变量
	root.Flags().StringVarP(&envFile, "envfile", "E", ".env", "设置ENV配置文件")

	if err := root.Execute(); err != nil {
		yiigo.Logger().Error("err cmd execute", zap.Error(err))
	}
}
