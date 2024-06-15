package identity

import (
	"context"
	"crypto/aes"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"api/lib/log"

	"github.com/pkg/errors"
	"github.com/shenghui0779/yiigo/xcrypto"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type CtxKeyAuth int

const IdentityKey CtxKeyAuth = 0

// Identity 授权身份
type Identity interface {
	// ID 授权ID
	ID() int64
	// Token 登录Token
	Token() string
	// String 用于日志记录
	String() string
	// AsAuthToken 生成授权Token
	AsAuthToken() (string, error)
}

type identity struct {
	I int64  `json:"i,omitempty"`
	T string `json:"t,omitempty"`
}

func (i *identity) ID() int64 {
	return i.I
}

func (i *identity) Token() string {
	return i.T
}

func (i *identity) String() string {
	if i.I == 0 {
		return "<nil>"
	}
	return fmt.Sprintf("id:%d|token:%s", i.I, i.T)
}

func (i *identity) AsAuthToken() (string, error) {
	b, err := json.Marshal(i)
	if err != nil {
		return "", errors.Wrap(err, "marshal identity")
	}

	key := []byte(viper.GetString("app.secret"))
	ct, err := xcrypto.AESEncryptCBC(key, key[:aes.BlockSize], b)
	if err != nil {
		return "", errors.Wrap(err, "encrypt identity")
	}
	return ct.String(), nil
}

// NewEmpty 空授权信息
func NewEmpty() Identity {
	return new(identity)
}

// New 用户授权信息
func New(id int64, token string) Identity {
	return &identity{
		I: id,
		T: token,
	}
}

// FromContext 获取授权信息
func FromContext(ctx context.Context) Identity {
	if ctx == nil {
		return NewEmpty()
	}

	identity, ok := ctx.Value(IdentityKey).(Identity)
	if !ok {
		return NewEmpty()
	}
	return identity
}

// FromAuthToken 解析授权Token
func FromAuthToken(ctx context.Context, token string) Identity {
	cipherText, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		log.Error(ctx, "Error invalid auth_token", zap.Error(err))
		return NewEmpty()
	}

	key := []byte(viper.GetString("app.secret"))
	plainText, err := xcrypto.AESDecryptCBC(key, key[:aes.BlockSize], cipherText)
	if err != nil {
		log.Error(ctx, "Error invalid auth_token", zap.Error(err))
		return NewEmpty()
	}

	identity := NewEmpty()
	if err = json.Unmarshal(plainText, identity); err != nil {
		log.Error(ctx, "Error invalid auth_token", zap.Error(err))
		// 此处应返回空Identify，因为若仅部分字段解析失败，Identity可能依然有效
		return NewEmpty()
	}
	return identity
}
