package base

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/pkg/errors"
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

type Client interface {
	Get(ctx context.Context, path string, query url.Values, resp Response) error
	Post(ctx context.Context, path string, body PostBody, resp Response) error
}

type client struct {
	host    string
	timeout time.Duration
}

func (c *client) Get(ctx context.Context, path string, query url.Values, resp Response) error {
	now := time.Now().Local()
	reqURL := c.url(path, query)
	logFields := make([]zap.Field, 0, 3)

	defer func() {
		logFields = append(logFields, zap.String("duration", time.Since(now).String()))

		yiigo.Logger().Info(fmt.Sprintf("[%s] [Iao.GET] %s", middleware.GetReqID(ctx), reqURL), logFields...)
	}()

	r, err := yiigo.HTTPGet(ctx, reqURL)

	if err != nil {
		logFields = append(logFields, zap.Error(err))

		return errors.Wrap(err, "Iao.Get error")
	}

	defer r.Body.Close()

	if r.StatusCode >= http.StatusBadRequest {
		io.Copy(ioutil.Discard, r.Body)

		err = fmt.Errorf("unexpected http status %d", r.StatusCode)

		logFields = append(logFields, zap.Error(err))

		return errors.Wrap(err, "Iao.Get error")
	}

	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		logFields = append(logFields, zap.Error(err))

		return errors.Wrap(err, "Iao.Get error")
	}

	logFields = append(logFields, zap.ByteString("response", b))

	if err = json.Unmarshal(b, resp); err != nil {
		logFields = append(logFields, zap.Error(err))

		return errors.Wrap(err, "Iao.Get error")
	}

	return nil
}

func (c *client) Post(ctx context.Context, path string, body PostBody, resp Response) error {
	now := time.Now().Local()
	reqURL := c.url(path, nil)
	logFields := make([]zap.Field, 0, 4)

	defer func() {
		logFields = append(logFields, zap.String("duration", time.Since(now).String()))

		yiigo.Logger().Info(fmt.Sprintf("[%s] [Iao.POST] %s", middleware.GetReqID(ctx), reqURL), logFields...)
	}()

	reqBody, err := body.Bytes()

	if err != nil {
		logFields = append(logFields, zap.Error(err))

		return errors.Wrap(err, "Iao.Post error")
	}

	logFields = append(logFields, zap.ByteString("body", reqBody))

	r, err := yiigo.HTTPPost(ctx, reqURL, reqBody, yiigo.WithHTTPHeader("Content-Type", "application/json; charset=utf-8"))

	if err != nil {
		logFields = append(logFields, zap.Error(err))

		return errors.Wrap(err, "Iao.Post error")
	}

	defer r.Body.Close()

	if r.StatusCode >= http.StatusBadRequest {
		io.Copy(ioutil.Discard, r.Body)

		err = fmt.Errorf("unexpected http status %d", r.StatusCode)

		logFields = append(logFields, zap.Error(err))

		return errors.Wrap(err, "Iao.Post error")
	}

	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		logFields = append(logFields, zap.Error(err))

		return errors.Wrap(err, "Iao.Post error")
	}

	logFields = append(logFields, zap.ByteString("response", b))

	if err = json.Unmarshal(b, resp); err != nil {
		logFields = append(logFields, zap.Error(err))

		return errors.Wrap(err, "Iao.Post error")
	}

	return nil
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
