package internal

import (
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRoundRobinReverseProxy_Empty(t *testing.T) {
	proxy := NewRoundRobinReverseProxy([]Backend{})
	assert.Empty(t, proxy.backends)
	assert.Equal(t, uint64(0), proxy.tcounter)
}

func TestRoundRobinReverseProxy_WithProxy(t *testing.T) {
	u1, _ := url.Parse("http://localhost:8081")
	backends := []Backend{{URL: *u1}}
	proxy := NewRoundRobinReverseProxy(backends)
	rp := proxy.WithProxy()
	assert.NotNil(t, rp)

	req := httptest.NewRequest("GET", "http://test", nil)
	rp.Director(req)
	assert.Equal(t, u1.Host, req.URL.Host)
	assert.Equal(t, u1.Scheme, req.URL.Scheme)
	assert.Equal(t, u1.Host, req.Host)
}

func TestNewRoundRobinReverseProxy_WithBackends(t *testing.T) {
	u1, _ := url.Parse("http://localhost:8081")
	u2, _ := url.Parse("http://localhost:8082")
	backends := []Backend{{URL: *u1}, {URL: *u2}}
	proxy := NewRoundRobinReverseProxy(backends)
	assert.Len(t, proxy.backends, 2)
	assert.Equal(t, *u1, proxy.backends[0])
	assert.Equal(t, *u2, proxy.backends[1])
}

func TestRoundRobinReverseProxy_nextBE(t *testing.T) {
	u1, _ := url.Parse("http://localhost:8081")
	u2, _ := url.Parse("http://localhost:8082")
	proxy := NewRoundRobinReverseProxy([]Backend{{URL: *u1}, {URL: *u2}})

	be1, err1 := proxy.nextBE()
	be2, err2 := proxy.nextBE()
	be3, err3 := proxy.nextBE()

	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.NoError(t, err3)
	assert.Equal(t, *u2, be1)
	assert.Equal(t, *u1, be2)
	assert.Equal(t, *u2, be3)
}

func TestRoundRobinReverseProxy_nextBE_NoBackends(t *testing.T) {
	proxy := NewRoundRobinReverseProxy([]Backend{})
	_, err := proxy.nextBE()
	assert.Error(t, err)
}

func TestRoundRobinReverseProxy_WithProxy_NoBackends(t *testing.T) {
	proxy := NewRoundRobinReverseProxy([]Backend{})
	rp := proxy.WithProxy()
	req := httptest.NewRequest("GET", "http://test", nil)
	rp.Director(req)
	assert.Nil(t, req.URL)
	assert.Equal(t, "backend unavailable", req.Header.Get("X-Reverse-Proxy-Error"))
}
