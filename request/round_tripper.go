package request

import (
	"net/http"
	"time"
)

type myRoundTripper struct {
	t 						http.RoundTripper
	maxRetry 			uint
	retryInterval time.Duration
}

func (r *myRoundTripper) RoundTrip(req *http.Request) (res *http.Response, err error) {
	for range(r.maxRetry + 1) {
		res, err = r.t.RoundTrip(req)
		if res == nil || err != nil { continue }
		if _, ok := retryableStatus[res.StatusCode]; !ok {
			return
		}

		select {
		case <-req.Context().Done():
			return nil, req.Context().Err()
		case <-time.After(r.retryInterval):
		}
	}
	return
}

var retryableStatus = map[int]struct{}{
	http.StatusInternalServerError: {},
	http.StatusTooManyRequests: {},
}