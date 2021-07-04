package base

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/pkg/errors"
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"

	"tplgo/internal/result"
)

type Client interface {
	Get(ctx context.Context, path string, query url.Values, resp Response) result.Result
	Post(ctx context.Context, path string, body PostBody, resp Response) result.Result
}

type client struct {
	host    string
	timeout time.Duration
}

func (c *client) Get(ctx context.Context, path string, query url.Values, resp Response) result.Result {
	now := time.Now().Local()
	reqURL := c.url(path, query)
	logFields := make([]zap.Field, 0, 4)

	defer func() {
		logFields = append(logFields, zap.String("request_id", middleware.GetReqID(ctx)))
		logFields = append(logFields, zap.String("duration", time.Since(now).String()))

		yiigo.Logger().Info(fmt.Sprintf("Iao.Get: %s", reqURL), logFields...)
	}()

	b, err := yiigo.HTTPGet(ctx, reqURL, yiigo.WithHTTPTimeout(c.timeout))

	if err != nil {
		logFields = append(logFields, zap.Error(err))

		return result.ErrIao.Wrap(result.WithErr(errors.Wrap(err, "Iao.Get error")))
	}

	logFields = append(logFields, zap.ByteString("response", b))

	if err = json.Unmarshal(b, resp); err != nil {
		logFields = append(logFields, zap.Error(err))

		return result.ErrIao.Wrap(result.WithErr(errors.Wrap(err, "Iao.Get error")))
	}

	return resp.ErrResult()
}

func (c *client) Post(ctx context.Context, path string, body PostBody, resp Response) result.Result {
	now := time.Now().Local()
	reqURL := c.url(path, nil)
	logFields := make([]zap.Field, 0, 5)

	defer func() {
		logFields = append(logFields, zap.String("request_id", middleware.GetReqID(ctx)))
		logFields = append(logFields, zap.String("duration", time.Since(now).String()))

		yiigo.Logger().Info(fmt.Sprintf("Iao.Post: %s", reqURL), logFields...)
	}()

	reqBody, err := body.Bytes()

	if err != nil {
		logFields = append(logFields, zap.Error(err))

		return result.ErrIao.Wrap(result.WithErr(errors.Wrap(err, "Iao.Post error")))
	}

	logFields = append(logFields, zap.ByteString("body", reqBody))

	b, err := yiigo.HTTPPost(ctx, reqURL, reqBody, yiigo.WithHTTPHeader("Content-Type", "application/json; charset=utf-8"), yiigo.WithHTTPTimeout(c.timeout))

	if err != nil {
		logFields = append(logFields, zap.Error(err))

		return result.ErrIao.Wrap(result.WithErr(errors.Wrap(err, "Iao.Post error")))
	}

	logFields = append(logFields, zap.ByteString("response", b))

	if err = json.Unmarshal(b, resp); err != nil {
		logFields = append(logFields, zap.Error(err))

		return result.ErrIao.Wrap(result.WithErr(errors.Wrap(err, "Iao.Post error")))
	}

	return resp.ErrResult()
}

func (c *client) url(path string, query url.Values) string {
	var builder strings.Builder

	builder.WriteString(c.host)

	if len(path) > 0 && path[:1] != "/" {
		builder.WriteString("/")
	}

	builder.WriteString(path)

	if len(query) != 0 {
		builder.WriteString("?")
		builder.WriteString(query.Encode())
	}

	return builder.String()
}

// NewClient returns new ms client
func NewClient(host string, timeout ...time.Duration) Client {
	if l := len(host); l > 0 && host[l-1:] == "/" {
		host = host[:l-1]
	}

	client := &client{
		host:    host,
		timeout: 5 * time.Second,
	}

	if len(timeout) > 0 {
		client.timeout = timeout[0]
	}

	return client
}
