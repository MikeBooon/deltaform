package middleware

import "github.com/labstack/echo/v4/middleware"

const (
	SecureRateLimitPerSecond = 3
)

var SecureRateLimitStore = middleware.NewRateLimiterMemoryStore(SecureRateLimitPerSecond)
