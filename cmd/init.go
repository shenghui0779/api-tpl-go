package cmd

import (
	"api/lib/log"

	"context"

	"go.uber.org/zap"
)

func Init() {
	// 注册命令
	rootCmd.AddCommand(helloCmd)
	rootCmd.AddCommand(migrateCmd)

	// 注册变量
	rootCmd.Flags().StringVarP(&cfgFile, "config", "C", "config.toml", "设置配置文件")

	if err := rootCmd.Execute(); err != nil {
		log.Error(context.Background(), "Error cmd execute", zap.Error(err))
	}
}
