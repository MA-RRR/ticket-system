package middleware

import (
	"net/http"
	"strings"

	"ticket-system/internal/pkg/errcode"
	"ticket-system/internal/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": errcode.ErrUnathorized.Error()})
			return
		}
		parts := strings.Fields(auth)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": errcode.ErrUnathorized.Error()})
			return
		}
		token := parts[1]
		claims, err := jwt.ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": errcode.ErrUnathorized.Error()})
			return
		}
		c.Set("user_id", claims.UID)
		c.Set("user_role", claims.Role)
		c.Next()
	}
}

func RequireRoles(roles ...string) gin.HandlerFunc {
	allowed := make(map[string]struct{}, len(roles))
	for _, r := range roles {
		allowed[r] = struct{}{}
	}
	return func(c *gin.Context) {
		v, ok := c.Get("user_role")
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": errcode.ErrForbidden.Error()})
			return
		}
		role, ok := v.(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": errcode.ErrForbidden.Error()})
			return
		}
		if _, ok := allowed[role]; !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": errcode.ErrForbidden.Error()})
			return
		}
		c.Next()
	}
}
