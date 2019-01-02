package controllers

import (
	"time"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/zuoyong8/coins/models"
)

var identityKey = "id"
type User struct {
	UserId    uint
	UserName  string
}

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func JwtAuth()(*jwt.GinJWTMiddleware, error){
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "chains_api",
		Key:         []byte("blockchain11668"),
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
			user,err := models.GetUsersByUsername(loginVals.Username)
			if err == nil {
				pwd,err := models.GetRealPwd(user.Pwdsalt,user.Password)
				if err == nil{
					if (loginVals.Username == user.Username  && loginVals.Password == pwd) {
						return &User{
							UserId:    user.Id,
							UserName:  user.Username,
						}, nil
					}
				}
			}
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			jwtClaims := jwt.ExtractClaims(c)
			if v, ok := data.(*User); ok && v.UserName == jwtClaims["id"].(string)  {
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

