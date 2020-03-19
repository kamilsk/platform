package config

import "time"

// HTTPServerConfig contains configuration for a common HTTP server.
// Deprecated: use go.octolab.org/toolkit/config/http.Server instead.
type HTTPServerConfig struct {
	Address           string        `json:"address"             yaml:"address"`
	ReadTimeout       time.Duration `json:"read-timeout"        yaml:"read-timeout"`
	ReadHeaderTimeout time.Duration `json:"read-header-timeout" yaml:"read-header-timeout"`
	WriteTimeout      time.Duration `json:"write-timeout"       yaml:"write-timeout"`
	IdleTimeout       time.Duration `json:"idle-timeout"        yaml:"idle-timeout"`
	MaxHeaderBytes    int           `json:"max-header-bytes"    yaml:"max-header-bytes"`
}
