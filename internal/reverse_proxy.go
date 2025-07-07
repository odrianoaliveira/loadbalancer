package internal

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

type RoundRobinReverseProxy struct {
	backends []url.URL
	tcounter uint64
}

func NewRoundRobinReverseProxy(backends []Backend) RoundRobinReverseProxy {
	var bes []url.URL

	for _, b := range backends {
		bes = append(bes, b.URL)
	}

	return RoundRobinReverseProxy{
		backends: bes,
		tcounter: 0,
	}
}

func (r *RoundRobinReverseProxy) WithProxy() *httputil.ReverseProxy {
	director := func(req *http.Request) {
		target, err := r.nextBE()
		if err != nil {
			slog.Error("reverse proxy error", "error", err)
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
