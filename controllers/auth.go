package controllers

import (
	"time"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

var identityKey = "id"
type User struct {
	UserId    int
	UserName  string
}

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}


func JwtAuth(username string,pwd string )(*jwt.GinJWTMiddleware, error){
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "chains_api",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				UserName: claims["id"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			userID := 1
			userName := loginVals.Username
			passWord := loginVals.Password

			if (userName == username && passWord == pwd) {
				return &User{
					UserId:    userID,
					UserName:  userName,
				}, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok && v.UserName == "admin" {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc: time.Now,
	})
	return authMiddleware,err
}

