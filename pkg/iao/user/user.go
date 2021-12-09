package user

import (
	"encoding/json"
	"tplgo/pkg/iao/client"
)

type ParamsUserInfo struct {
	UserIDs []int64  `json:"user_ids"`
	Columns []string `json:"columns"`
}

type ResultUserInfo struct {
	ID           int64  `json:"id"`
	Nickname     string `json:"nickname"`
	Avatar       string `json:"avatar"`
	AuthID       string `json:"auth_id"`
	Secret       string `json:"secret"`
	VIPLevel     int    `json:"vip_level"`
	VIPExpiredAt int64  `json:"vip_expired_at"`
}

func GetUserInfo(params *ParamsUserInfo, result *ResultUserInfo) client.Action {
	return client.NewPostAction("/users/info",
		client.WithBody(func() ([]byte, error) {
			return json.Marshal(params)
		}),
		client.WithDecode(func(b []byte) error {
			resp := client.NewResponse(result)

			if err := json.Unmarshal(b, resp); err != nil {
				return err
			}

			return resp.Error()
		}),
	)
}
