package memory

import (
	"time"

	"github.com/fasthttp/session/v2"
)

// Config configuration of provider
type Config struct{}

// Provider backend manager
type Provider struct {
	config Config
	db     *session.Dict
}

type item struct {
	data           []byte
	lastActiveTime int64
	expiration     time.Duration
}