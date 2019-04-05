package http

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	cacheControl = "Cache-Control"
	noCache      = "no-cache"
	xStrict      = "X-Strict"
	xTimeout     = "X-Timeout"
)

func Timeout(header http.Header, fallback time.Duration, percentage float64) time.Duration {
	d, err := time.ParseDuration(header.Get(xTimeout))
	if err != nil {
		d = fallback
	}
	return time.Duration(percentage * float64(d))
}

func NoCache(header http.Header) bool {
	return strings.EqualFold(header.Get(cacheControl), noCache)
}

func Strict(header http.Header) bool {
	if v := header.Get(xStrict); v != "" {
		isStrict, _ := strconv.ParseBool(v)
		return isStrict
	}
	return false
}
