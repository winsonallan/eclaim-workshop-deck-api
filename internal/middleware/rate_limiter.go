package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	limitergin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

func RateLimiter(rateStr string) (gin.HandlerFunc, error) {
	// Format: "X-period" e.g. "100-M" = 100 req/min, "1000-H" = 1000 req/hour
	rate, err := limiter.NewRateFromFormatted(rateStr)
	if err != nil {
		return nil, err
	}

	store := memory.NewStore()
	instance := limiter.New(store, rate)

	middleware := limitergin.NewMiddleware(instance, limitergin.WithLimitReachedHandler(func(c *gin.Context) {
		c.JSON(http.StatusTooManyRequests, gin.H{
			"success": false,
			"message": "too many requests, slow down",
		})
		c.Abort()
	}))

	return middleware, nil
}
