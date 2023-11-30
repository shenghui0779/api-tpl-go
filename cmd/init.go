package cmd

import (
	"api/lib/log"

	"context"

	"go.uber.org/zap"
)

func Init() {
	// 注册命令
	rootCmd.AddCommand(helloCmd)

	// 注册变量
	rootCmd.Flags().StringVarP(&cfgFile, "config", "C", ".yml", "设置配置文件")

	if err := rootCmd.Execute(); err != nil {
		log.Error(context.Background(), "err cmd execute", zap.Error(err))
	}
}
