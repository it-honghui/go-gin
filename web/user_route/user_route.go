package user_route

import (
	"github.com/gin-gonic/gin"
	"go-gin/domain"
	"go-gin/domain/entity"
	"go-gin/domain/enum"
	"go-gin/middleware"
	"go-gin/service/user_service"
	"strconv"
)

func login(c *gin.Context) {
	token := user_service.Login(c.PostForm("username"), c.PostForm("password"))
	domain.Ok(c, gin.H{
		"token": token,
	})
}

func createUser(c *gin.Context) {
	var user entity.User
	_ = c.ShouldBindJSON(&user)
	user, err := user_service.CreateUser(&user)
	if err != nil {
		domain.Panic(domain.USERNAME_ALREADY_EXIST, "")
	}
	domain.Ok(c, user)
}

func findUsers() gin.HandlerFunc {
	return middleware.Authenticator(func(c *gin.Context, user *entity.User) {
		if user.Role == enum.ADMIN {
			users, err := user_service.FindUsers()
			if err != nil {
				domain.Panic(domain.ERROR, err.Error())
			}
			domain.Ok(c, users)
		} else {
			domain.Unauthorized(c, domain.USER_UNAUTHORIZED.Code, domain.USER_UNAUTHORIZED.Msg)
		}
	})
}

func findUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := user_service.FindUser(uint64(id))
	if err != nil {
		domain.Panic(domain.NOT_FOUND, "user")
	}
	domain.Ok(c, user)
}

func deleteUser() gin.HandlerFunc {
	return middleware.Authenticator(func(c *gin.Context, user *entity.User) {
		if user.Role == enum.ADMIN {
			id, _ := strconv.Atoi(c.Param("id"))
			err := user_service.DeleteUser(uint64(id))
			if err != nil {
				domain.Panic(domain.ERROR, err.Error())
			}
		} else {
			domain.Unauthorized(c, domain.USER_UNAUTHORIZED.Code, domain.USER_UNAUTHORIZED.Msg)
		}
	})
}

func updatePassword(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := user_service.UpdatePassword(uint64(id), c.PostForm("password"))
	if err != nil {
		domain.Panic(domain.ERROR, err.Error())
	}
}

func Setup(e *gin.Engine) {
	g := e.Group("/user")
	{
		g.POST("/login", login)
		g.POST("/", createUser)
		g.GET("/", middleware.JWTAuthenticatorMiddleware(), findUsers())
		g.GET("/:id", middleware.JWTAuthenticatorMiddleware(), findUser)
		g.DELETE("/:id", middleware.JWTAuthenticatorMiddleware(), deleteUser())
		g.PATCH("/:id/password", middleware.JWTAuthenticatorMiddleware(), updatePassword)
	}
}
