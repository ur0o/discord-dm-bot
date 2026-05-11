package request

import (
	"time"
)

type options struct {
	MaxRetry 			int
	RetryInterval time.Duration
	Timeout 			time.Duration
}

type Option func(o *options)

func MaxRetry(maxRetry int) Option {
	return func(opt *options) {
		opt.MaxRetry = maxRetry
	}
}

func RetryInterval(retryInterval time.Duration) Option {
	return func(opt *options) {
		opt.RetryInterval = retryInterval
	}
}

func Timeout(timeout time.Duration) Option {
	return func(opt *options) {
		opt.Timeout = timeout
	}
}