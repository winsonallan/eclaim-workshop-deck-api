package middleware

import (
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
)

func SecurityHeaders(env string) gin.HandlerFunc {
	return secure.New(secure.Config{
		// Forces HTTPS — only enable in production (will break local dev)
		SSLRedirect: env == "production",

		// HSTS: once visited over HTTPS, browser won't allow HTTP for 1 year
		STSSeconds:           31536000,
		STSIncludeSubdomains: true,

		// Prevent clickjacking — blocks your app from being iframed
		FrameDeny: true,

		// Don't let browsers guess the content type
		ContentTypeNosniff: true,

		// Block XSS in older browsers
		BrowserXssFilter: true,

		// Referrer policy — controls how much referrer info is sent
		ReferrerPolicy: "strict-origin-when-cross-origin",

		// Content Security Policy
		// Adjust based on what your API actually serves
		ContentSecurityPolicy: "default-src 'self'",

		// Prevent IE from switching to compatibility mode
		IENoOpen: true,
	})
}
