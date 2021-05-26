package user

type ParamsUserInfo struct {
	UserIDs []int64  `json:"user_ids"`
	Columns []string `json:"columns"`
}
