package command

import (
	"github.com/gsxhnd/garage/crawl"
	"github.com/gsxhnd/garage/utils"
	"github.com/urfave/cli/v2"
)

var codeCmd = &cli.Command{
	Name:      "code",
	Aliases:   nil,
	Usage:     "根据指定番号爬取数据",
	UsageText: "code --site [javbus/javlibrary] XXX-001",
	Flags: []cli.Flag{
		proxyFlag,
		siteFlag,
	},
	Action: func(ctx *cli.Context) error {
		newLogger := utils.GetLogger()
		var code = ctx.Args().Get(0)
		c := crawl.NewCrawlClient(newLogger)
		_ = c.SetProxy(ctx.String("proxy"))
		c.StarCrawlJavbusMovie(code)
		return nil
	},
}
