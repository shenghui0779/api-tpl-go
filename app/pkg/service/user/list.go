package user

import (
	"context"

	"api/app/ent"
	"api/app/ent/user"
	"api/lib"
	"api/lib/log"
	"api/lib/result"

	"github.com/tidwall/gjson"
	"go.uber.org/zap"
)

type ReqList struct {
	Page   int      `json:"page"`
	Size   int      `json:"size"`
	Query  string   `json:"query"` // `{"account":"xxx"}`
	Fields []string `json:"fields"`
}

type RespList struct {
	Total int         `json:"total"`
	List  []*UserInfo `json:"list"`
}

func List(ctx context.Context, req *ReqList) result.Result {
	query := gjson.Parse(req.Query)
	builder := ent.DB.User.Query()
	if v := query.Get("account").String(); len(v) != 0 {
		builder.Where(user.Account(v))
	}

	total := 0
	offset, limit := lib.PostPage(req.Page, req.Size)

	// 仅第一页返回数量并由前端保存
	if offset == 0 {
		var err error
		total, err = builder.Clone().Unique(false).Count(ctx)
		if err != nil {
			log.Error(ctx, "Error ent.User.Query.Count", zap.Error(err))
			return result.ErrSystem(result.E(err))
		}
	}

	records, err := builder.Unique(false).Order(ent.Desc(user.FieldID)).Offset(offset).Limit(limit).All(ctx)
	if err != nil {
		log.Error(ctx, "Error ent.User.Query.All", zap.Error(err))
		return result.ErrSystem(result.E(err))
	}

	resp := &RespList{
		Total: total,
		List:  make([]*UserInfo, 0, len(records)),
	}
	for _, v := range records {
		resp.List = append(resp.List, EntToUserInfo(v))
	}
	return result.OK(result.V(resp))
}
