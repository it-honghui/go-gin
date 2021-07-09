package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-gin/config"
	"go-gin/domain"
	"go-gin/domain/entity"
	"strings"
	"time"
)

var SecretKey = config.Config.JWT.SecretKey

func GenerateJWT(username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	// 使用一个密钥字符串对 Token 进行签名
	t, _ := token.SignedString([]byte(SecretKey))
	return t
}

func JWTAuthenticatorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 Authorization 提取 JWT
		jwtStr := c.Request.Header.Get("Authorization")
		// 校验 JWT
		if jwtStr == "" {
			domain.Unauthorized(c, domain.TOKEN_NOT_EXIST.Code, domain.TOKEN_NOT_EXIST.Msg)
			c.Abort()
			return
		}
		checkToken := strings.Split(jwtStr, " ")
		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			domain.Unauthorized(c, domain.TOKEN_TYPE_ERROR.Code, domain.TOKEN_TYPE_ERROR.Msg)
			c.Abort()
			return
		}
		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(checkToken[1], claims, func(*jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					domain.Unauthorized(c, domain.TOKEN_EXPIRE.Code, domain.TOKEN_EXPIRE.Msg)
					c.Abort()
					return
				} else {
					domain.Unauthorized(c, domain.TOKEN_ERROR.Code, domain.TOKEN_ERROR.Msg)
					c.Abort()
					return
				}
			}
		}
		user := entity.User{}
		config.DB.Where("username = ?", claims["username"].(string)).Omit("password").First(&user)
		if user.ID == 0 {
			domain.Unauthorized(c, domain.NOT_FOUND.Code, fmt.Sprintf("%s : %s", domain.NOT_FOUND, claims["username"].(string)))
			c.Abort()
			return
		}
		c.Set("user", user)
	}
}

type AuthAPIFunc func(c *gin.Context, user *entity.User)

func Authenticator(f AuthAPIFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		u, exists := c.Get("user")
		if !exists {
			domain.Unauthorized(c, domain.TOKEN_NOT_EXIST.Code, domain.TOKEN_NOT_EXIST.Msg)
			c.Abort()
			return
		}
		var user entity.User
		user = u.(entity.User)
		f(c, &user)
	}
}
