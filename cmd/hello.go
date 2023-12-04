package cmd

import (
	"api/pkg/script"

	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:     "hello",
	Short:   "命令示例",
	Long:    "命令示例",
	Version: "v1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		script.Hello()
	},
}
