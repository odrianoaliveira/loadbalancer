package internal

type Backend struct {
	URL         string
	IsAlive     bool
	connections int
}
