package lib

import (
	"context"
	"crypto/aes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"time"

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
	// Token 授权Token
	Token() string
	// Encrypt 授权加密
	Encrypt() (string, error)
	// Decrypt 授权解密
	Decrypt(cipherText []byte) error
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

func (i *identity) Encrypt() (string, error) {
	plainText, err := json.Marshal(i)

	if err != nil {
		return "", errors.Wrap(err, "marshal identity")
	}

	key := []byte(os.Getenv("API_SECRET"))
	iv := key[:aes.BlockSize]

	cryptor := yiigo.NewCBCCrypto(key, iv, yiigo.PKCS7)

	cipherText, err := cryptor.Encrypt(plainText)

	if err != nil {
		return "", errors.Wrap(err, "encrypt identity")
	}

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func (i *identity) Decrypt(cipherText []byte) error {
	key := []byte(os.Getenv("API_SECRET"))
	iv := key[:aes.BlockSize]

	cryptor := yiigo.NewCBCCrypto(key, iv, yiigo.PKCS7)

	plainText, err := cryptor.Decrypt(cipherText)

	if err != nil {
		return errors.Wrap(err, "decrypt identity")
	}

	if err = json.Unmarshal(plainText, i); err != nil {
		return errors.Wrap(err, "unmarshal identity")
	}

	return nil
}

// NewEmptyIdentity 空授权信息
func NewEmptyIdentity() Identity {
	return new(identity)
}

// NewIdentity 用户授权信息
func NewIdentity(userID int64) Identity {
	return &identity{
		I: userID,
		T: yiigo.MD5(fmt.Sprintf("%d.%d.%s", userID, time.Now().Unix(), Nonce())),
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

func VerifyAuthToken(ctx context.Context, token string) (Identity, error) {
	cipherText, err := base64.StdEncoding.DecodeString(token)

	if err != nil {
		logger.Err(ctx, "err auth (base64_decode)", zap.Error(err))

		return nil, errors.New("授权信息错误，请重新登录")
	}

	identity := NewEmptyIdentity()

	if err := identity.Decrypt(cipherText); err != nil {
		logger.Err(ctx, "err auth (decrypt)", zap.Error(err))

		return nil, errors.New("授权信息错误，请重新登录")
	}

	if identity.ID() == 0 {
		return nil, errors.New("未授权，请先登录")
	}

	record, err := ent.DB.User.Query().Unique(false).Select(user.FieldID, user.FieldLoginToken).Where(user.ID(identity.ID())).First(ctx)

	if err != nil {
		logger.Err(ctx, "err auth (query user)", zap.Error(err))

		return nil, errors.New("内部服务器错误")
	}

	if len(record.LoginToken) == 0 || record.LoginToken != identity.Token() {
		return nil, errors.New("授权已失效")
	}

	return identity, nil
}
