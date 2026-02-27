package client

import (
	"crypto/tls"
	"time"

	"github.com/go-resty/resty/v2"
	"errors"
)

// Option 客户端配置选项函数（Functional Options模式）
type Option func(*Client) error

// WithTimeout 设置请求超时时间
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) error {
		if timeout <= 0 {
			return errors.New("timeout must be greater than 0")
		}
		c.restyClient.SetTimeout(timeout)
		c.config.Timeout = timeout
		return nil
	}
}

// WithHeaders 设置全局请求头
func WithHeaders(headers map[string]string) Option {
	return func(c *Client) error {
		if headers == nil {
			return errors.New("headers cannot be nil")
		}
		c.restyClient.SetHeaders(headers)
		return nil
	}
}

// WithRetry 设置重试机制
func WithRetry(maxRetries int, retryWaitTime time.Duration) Option {
	return func(c *Client) error {
		if maxRetries < 0 {
			return errors.New("maxRetries must be >= 0")
		}
		if retryWaitTime < 0 {
			return errors.New("retryWaitTime must be >= 0")
		}
		c.restyClient.
			SetRetryCount(maxRetries).
			SetRetryWaitTime(retryWaitTime)
		return nil
	}
}

// WithDebugMode 启用调试模式（打印请求/响应详情）
func WithDebugMode(enable bool) Option {
	return func(c *Client) error {
		c.restyClient.SetDebug(enable)
		return nil
	}
}

// WithProxy 设置代理
func WithProxy(proxyURL string) Option {
	return func(c *Client) error {
		if proxyURL == "" {
			return errors.New("proxyURL cannot be empty")
		}
		c.restyClient = c.restyClient.SetProxy(proxyURL)
		return nil
	}
}

// WithTLSConfig 自定义TLS配置（如跳过证书验证）
func WithTLSConfig(insecureSkipVerify bool) Option {
	return func(c *Client) error {
		c.restyClient.SetTLSClientConfig(&tls.Config{
			InsecureSkipVerify: insecureSkipVerify,
		})
		return nil
	}
}