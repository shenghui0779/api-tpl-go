package scripts

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdHello = &cobra.Command{
	Use:     "hello",
	Short:   "命令示例",
	Long:    "命令示例",
	Version: "v1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello world")
	},
}
