package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func extractUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetHeader("X-User-ID")
		if userID != "" {
			if _, err := uuid.Parse(userID); err != nil {
				c.AbortWithStatusJSON(401, gin.H{"error": 401, "message": "Invalid User ID"})
				return
			}
		}
		userType := c.GetHeader("X-User-Type")
		logrus.Infof("UserID :: %v, UserType :: %v", userID, userType)
		c.Set("UserID", userID)
		c.Set("UserType", userType)
		c.Next()
	}
}
