package media

import (
	"time"

	"github.com/shenghui0779/yiigo"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"golang.org/x/net/context"

	"api/app/ent"
	"api/app/ent/media"
	"api/lib"
	"api/lib/log"
	"api/lib/result"
)

type ReqList struct {
	Page  int    `json:"page"`
	Size  int    `json:"size"`
	Query string `json:"query"` // `{"media_id":"","file_name":""}`
}

type RespList struct {
	Total int         `json:"total"`
	List  []MediaInfo `json:"list"`
}

func List(ctx context.Context, req *ReqList) result.Result {
	query := gjson.Parse(req.Query)
	builder := ent.DB.Media.Query()
	if v := query.Get("media_id").String(); len(v) != 0 {
		builder.Where(media.MediaID(v))
	}
	if v := query.Get("file_name").String(); len(v) != 0 {
		builder.Where(media.FileNameHasPrefix(v))
	}

	total := 0
	offset, limit := lib.PostPage(req.Page, req.Size)

	// 仅第一页返回数量并由前端保存
	if offset == 0 {
		var err error

		total, err = builder.Clone().Unique(false).Count(ctx)
		if err != nil {
			log.Error(ctx, "Error count account", zap.Error(err))
			return result.ErrSystem(result.E(err))
		}
	}

	records, err := builder.WithFile(func(fq *ent.FileQuery) {
		fq.Unique(false)
	}).Unique(false).Order(ent.Desc(media.FieldID)).Offset(offset).Limit(limit).All(ctx)
	if err != nil {
		log.Error(ctx, "Error query media", zap.Error(err))
		return result.ErrSystem(result.E(err))
	}

	resp := &RespList{
		Total: total,
		List:  make([]MediaInfo, 0, len(records)),
	}

	for _, v := range records {
		info := MediaInfo{
			MediaID: v.MediaID,
			Name:    v.FileName,
			ModTime: v.CreatedAt.In(time.Local).Format(time.DateTime),
		}
		if ef := v.Edges.File; ef != nil {
			info.Size = ef.Size
			info.SizeStr = yiigo.Quantity(ef.Size).String()
			info.Format = ef.Format
			info.Width = ef.Width
			info.Height = ef.Height
			info.Orientation = Orientation(ef.Orientation).String()
			info.Duration = ef.Duration
			info.Fingerprint = ef.Fingerprint
		}
		resp.List = append(resp.List, info)
	}

	return result.OK(result.V(resp))
}
