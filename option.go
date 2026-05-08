package ddm

import (
	"time"
)

type requestOptions struct {
	MaxRetry 			int
	RetryInterval time.Duration
	Timeout 			time.Duration
}

type RequestOption func(o *requestOptions)

func MaxRetry(maxRetry int) RequestOption {
	return func(opt *requestOptions) {
		opt.MaxRetry = maxRetry
	}
}

func RetryInterval(retryInterval time.Duration) RequestOption {
	return func(opt *requestOptions) {
		opt.RetryInterval = retryInterval
	}
}

func Timeout(timeout time.Duration) RequestOption {
	return func(opt *requestOptions) {
		opt.Timeout = timeout
	}
}