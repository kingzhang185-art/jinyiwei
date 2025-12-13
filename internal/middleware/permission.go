package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// RequireRole 要求特定角色的中间件
func RequireRole(roleCodes ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取用户角色
		roles, exists := c.Get("roles")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "无权限访问",
			})
			c.Abort()
			return
		}

		userRoles, ok := roles.([]string)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "无权限访问",
			})
			c.Abort()
			return
		}

		// 检查是否有匹配的角色
		hasRole := false
		for _, userRole := range userRoles {
			for _, requiredRole := range roleCodes {
				if userRole == requiredRole {
					hasRole = true
					break
				}
			}
			if hasRole {
				break
			}
		}

		if !hasRole {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "无权限访问，需要角色: " + strings.Join(roleCodes, ", "),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

