package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/joho/godotenv/autoload"
	"github.com/xerdin442/api-practice/internal/cache"
	"github.com/xerdin442/api-practice/internal/env"
)

var secretKey = []byte(env.GetStr("JWT_SECRET"))

type AllClaims struct {
	UserID int32 `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userID int32) (string, error) {
	claims := AllClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func AuthMiddleware(cache *cache.Redis) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.ParseWithClaims(tokenString, &AllClaims{}, func(token *jwt.Token) (any, error) {
			return secretKey, nil
		})

		isBlacklisted, err := cache.IsBlacklisted(c.Request.Context(), tokenString)

		if err != nil || !token.Valid || isBlacklisted {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Session expired. Please log in"})
			return
		}

		claims, ok := token.Claims.(*AllClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("token_exp", time.Unix(claims.ExpiresAt.Unix(), 0))
		c.Next()
	}
}
