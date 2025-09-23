package helpers

// func RateLimiter() gin.HandlerFunc {
// 	var limiter = rate.NewLimiter(1, 5)
// 	return func(c *gin.Context) {

// 		if !limiter.Allow() {
// 			c.JSON(http.StatusTooManyRequests, gin.H{
// 				"error": "too many requests",
// 			})
// 			c.Abort()
// 		}
// 		c.Next()
// 	}
// }
