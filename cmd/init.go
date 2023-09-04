package cmd

import (
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

func Init() {
	// 注册命令
	rootCmd.AddCommand(helloCmd)

	// 注册变量
	rootCmd.Flags().StringVarP(&envFile, "envfile", "E", ".env", "设置ENV配置文件")

	if err := rootCmd.Execute(); err != nil {
		yiigo.Logger().Error("err cmd execute", zap.Error(err))
	}
}
