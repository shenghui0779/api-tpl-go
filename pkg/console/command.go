package console

import (
	"context"
	"tplgo/pkg/ent"
	"tplgo/pkg/ent/migrate"

	"github.com/shenghui0779/yiigo"
	"github.com/urfave/cli/v2"
)

var Commands = []*cli.Command{
	{
		Name:    "migrate",
		Aliases: []string{"M"},
		Usage:   "数据库迁移",
		Action: func(c *cli.Context) error {
			client := ent.NewClient(ent.Driver(yiigo.EntDriver()))
			defer client.Close()

			return client.Schema.Create(context.Background(),
				migrate.WithDropIndex(true),
				migrate.WithDropColumn(true),
				migrate.WithForeignKeys(false),
			)
		},
	},
}
