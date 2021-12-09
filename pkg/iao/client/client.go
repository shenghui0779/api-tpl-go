package client

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/shenghui0779/yiigo"
)

// Client is the interface that do http request
type Client interface {
	Do(ctx context.Context, action Action, options ...yiigo.HTTPOption) error
}

type apiclient struct {
	host   string
	client yiigo.HTTPClient
	logger Logger
}

func (c *apiclient) Do(ctx context.Context, action Action, options ...yiigo.HTTPOption) error {
	reqURL := action.URL(c.host)

	logData := &LogData{
		URL:    reqURL,
		Method: action.Method(),
	}

	now := time.Now().Local()

	defer func() {
		logData.Duration = time.Since(now)
		c.logger.Log(ctx, logData)
	}()

	body, err := action.Body()

	if err != nil {
		logData.Error = err

		return err
	}

	if len(body) != 0 {
		logData.Body = body
		options = append(options, yiigo.WithHTTPHeader("Content-Type", "application/json; charset=utf-8"))
	}

	resp, err := c.client.Do(ctx, action.Method(), reqURL, body, options...)

	if err != nil {
		logData.Error = err

		return err
	}

	defer resp.Body.Close()

	logData.StatusCode = resp.StatusCode

	if resp.StatusCode >= http.StatusBadRequest {
		io.Copy(ioutil.Discard, resp.Body)

		return fmt.Errorf("unexpected status %d", resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		logData.Error = err

		return err
	}

	logData.Response = b

	if err = action.Decode(b); err != nil {
		logData.Error = err

		return err
	}

	return nil
}

// NewClient returns a new wechat client
func NewClient(host string, c *http.Client) Client {
	return &apiclient{
		host:   host,
		client: yiigo.NewHTTPClient(c),
		logger: NewLogger(),
	}
}

// NewDefaultClient returns a new default wechat client
func NewDefaultClient(host string) Client {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 60 * time.Second,
			}).DialContext,
			MaxIdleConns:          0,
			MaxIdleConnsPerHost:   1000,
			MaxConnsPerHost:       1000,
			IdleConnTimeout:       60 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}

	return &apiclient{
		host:   host,
		client: yiigo.NewHTTPClient(client),
		logger: NewLogger(),
	}
}
