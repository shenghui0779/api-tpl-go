package media

import (
	"net/url"
	"time"

	"github.com/pkg/errors"
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
	"golang.org/x/net/context"

	"api/ent"
	"api/ent/media"
	"api/lib/db"
	"api/lib/log"
	"api/pkg/consts"
	"api/pkg/internal"
	"api/pkg/result"
)

type RespList struct {
	Total int         `json:"total"`
	List  []MediaInfo `json:"list"`
}

func List(ctx context.Context, query url.Values) result.Result {
	builder := db.Client().Media.Query()

	total := 0
	offset, limit := internal.QueryPage(ctx, query)

	// 仅第一页返回数量并由前端保存
	if offset == 0 {
		var err error

		total, err = builder.Clone().Unique(false).Count(ctx)
		if err != nil {
			log.Error(ctx, "Error count account", zap.Error(err))
			return result.ErrSystem(result.E(errors.WithMessage(err, "用户Count失败")))
		}
	}

	records, err := builder.WithFile(func(fq *ent.FileQuery) {
		fq.Unique(false)
	}).Unique(false).Order(ent.Desc(media.FieldID)).Offset(offset).Limit(limit).All(ctx)
	if err != nil {
		log.Error(ctx, "Error query media", zap.Error(err))
		return result.ErrSystem(result.E(errors.WithMessage(err, "用户查询失败")))
	}

	resp := &RespList{
		Total: total,
		List:  make([]MediaInfo, 0, len(records)),
	}

	for _, v := range records {
		info := MediaInfo{
			MediaID: v.MediaID,
			Name:    v.FileName,
			ModTime: v.CreatedAt.Format(time.DateTime),
		}
		if ef := v.Edges.File; ef != nil {
			info.Size = ef.Size
			info.SizeStr = yiigo.Quantity(ef.Size).String()
			info.Format = ef.Format
			info.Width = ef.Width
			info.Height = ef.Height
			info.Orientation = consts.Orientation(ef.Orientation).String()
			info.Fingerprint = ef.Fingerprint
		}
		resp.List = append(resp.List, info)
	}

	return result.OK(result.V(resp))
}
