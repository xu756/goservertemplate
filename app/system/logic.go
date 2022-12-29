package system

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"server/common/log"
	"server/common/model"
)

func Reset(ctx context.Context, c *app.RequestContext) {
	model.SqliteCreateTable("logs", log.Logs{})
}
