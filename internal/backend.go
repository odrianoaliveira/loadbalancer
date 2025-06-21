package internal

import "net/url"

type Backend struct {
	URL         *url.URL
	IsAlive     bool
	connections int
}
