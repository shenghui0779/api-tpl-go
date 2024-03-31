package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"api/ent/migrate"
	"api/lib/db"
	"api/lib/log"
)

var migrateCmd = &cobra.Command{
	Use:     "migrate",
	Short:   "迁移",
	Long:    "数据库迁移",
	Version: "v1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		if err := db.Client().Schema.Create(ctx,
			migrate.WithForeignKeys(false),
		); err != nil {
			log.Error(ctx, "数据库表迁移失败", zap.Error(err))
		}
	},
}
