package internal

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"

	"go.uber.org/zap"
)

type RoundRobinReverseProxy struct {
	logger   *zap.Logger
	backends []url.URL
	tcounter uint64
}

func NewRoundRobinReverseProxy(backends []Backend, log *zap.Logger) RoundRobinReverseProxy {
	var bes []url.URL

	for _, b := range backends {
		bes = append(bes, b.URL)
	}

	return RoundRobinReverseProxy{
		logger:   log,
		backends: bes,
		tcounter: 0,
	}
}

func (r *RoundRobinReverseProxy) WithProxy() *httputil.ReverseProxy {
	director := func(req *http.Request) {
		target, err := r.nextBE()
		if err != nil {
			r.logger.Error("reverse proxy error", zap.Error(err))
			req.URL = nil
			req.Header.Set("X-Reverse-Proxy-Error", "backend unavailable")
			return
		}
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.Host = target.Host
	}

	return &httputil.ReverseProxy{Director: director}
}

func (r *RoundRobinReverseProxy) nextBE() (url.URL, error) {
	if len(r.backends) == 0 {
		return url.URL{}, fmt.Errorf("no backends available")
	}
	atomic.AddUint64(&r.tcounter, 1)

	n := r.tcounter % uint64(len(r.backends))
	return r.backends[n], nil
}
