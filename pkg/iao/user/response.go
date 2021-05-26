package user

type UserInfo struct {
	ID           int64  `json:"id"`
	Nickname     string `json:"nickname"`
	Avatar       string `json:"avatar"`
	AuthID       string `json:"auth_id"`
	Secret       string `json:"secret"`
	VIPLevel     int    `json:"vip_level"`
	VIPExpiredAt int64  `json:"vip_expired_at"`
}
