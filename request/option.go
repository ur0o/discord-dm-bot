package request

import (
	"time"
)

type options struct {
	MaxRetry 			uint
	RetryInterval time.Duration
	Timeout 			time.Duration
}

type Option func(o *options)

func MaxRetry(maxRetry uint) Option {
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