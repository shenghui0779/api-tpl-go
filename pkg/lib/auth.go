package lib

import (
	"context"
	"crypto/aes"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"tplgo/pkg/config"
	"tplgo/pkg/ent"
	"tplgo/pkg/ent/user"
	"tplgo/pkg/logger"

	"github.com/pkg/errors"
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

type CtxKeyAuth int

const AuthIdentityKey CtxKeyAuth = 0

// Identity 授权身份
type Identity interface {
	// ID 授权ID
	ID() int64
	// Encrypt 授权加密
	Encrypt() (string, error)
	// Decrypt 授权解密
	Decrypt(cipherText []byte) error
	// Check 校验
	Check(ctx context.Context) error
	// String 用于日志记录
	String() string
}

type identity struct {
	I int64  `json:"i,omitempty"`
	T string `json:"t,omitempty"`
}

func (i *identity) ID() int64 {
	return i.I
}

func (i *identity) Encrypt() (string, error) {
	plainText, err := json.Marshal(i)

	if err != nil {
		return "", errors.Wrap(err, "marshal identity")
	}

	key := []byte(config.ENV.APISecret)
	iv := key[:aes.BlockSize]

	cryptor := yiigo.NewCBCCrypto(key, iv, yiigo.AES_PKCS5)

	cipherText, err := cryptor.Encrypt(plainText)

	if err != nil {
		return "", errors.Wrap(err, "encrypt identity")
	}

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func (i *identity) Decrypt(cipherText []byte) error {
	key := []byte(config.ENV.APISecret)
	iv := key[:aes.BlockSize]

	cryptor := yiigo.NewCBCCrypto(key, iv, yiigo.AES_PKCS5)

	plainText, err := cryptor.Decrypt(cipherText)

	if err != nil {
		return errors.Wrap(err, "decrypt identity")
	}

	if err = json.Unmarshal(plainText, i); err != nil {
		return errors.Wrap(err, "unmarshal identity")
	}

	return nil
}

func (i *identity) Check(ctx context.Context) error {
	if i.I == 0 {
		return errors.New("未授权，请先登录")
	}

	record, err := ent.DB.User.Query().Unique(false).Select(
		user.FieldID,
		user.FieldLoginToken,
	).Where(user.ID(i.I)).First(ctx)

	if err != nil {
		logger.Err(ctx, "err auth check", zap.Error(err))

		return errors.New("内部服务器错误")
	}

	if len(record.LoginToken) == 0 || record.LoginToken != i.T {
		return errors.New("授权已失效")
	}

	return nil
}

func (i *identity) String() string {
	if i.I == 0 {
		return ""
	}

	return fmt.Sprintf("id:%d|token:%s", i.I, i.T)
}

// NewEmptyIdentity 空授权信息
func NewEmptyIdentity() Identity {
	return new(identity)
}

// NewIdentity 用户授权信息
func NewIdentity(id int64, token string) Identity {
	return &identity{
		I: id,
		T: token,
	}
}

// GetIdentity 获取授权信息
func GetIdentity(ctx context.Context) Identity {
	if ctx == nil {
		return NewEmptyIdentity()
	}

	identity, ok := ctx.Value(AuthIdentityKey).(Identity)

	if !ok {
		return NewEmptyIdentity()
	}

	return identity
}

// AuthTokenToIdentity 解析授权Token
func AuthTokenToIdentity(ctx context.Context, token string) Identity {
	cipherText, err := base64.StdEncoding.DecodeString(token)

	if err != nil {
		logger.Err(ctx, "err invalid auth_token", zap.Error(err))

		return NewEmptyIdentity()
	}

	identity := NewEmptyIdentity()

	if err := identity.Decrypt(cipherText); err != nil {
		logger.Err(ctx, "err invalid auth_token", zap.Error(err))

		return NewEmptyIdentity()
	}

	return identity
}
