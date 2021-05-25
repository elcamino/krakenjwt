package krakenjwt

import (
	"encoding/hex"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const APISecret = "5fc72dcca1881e421a6d92b2725d8cf6315e1c6f7657e7ba7e3cd6366fb71066"
const identityKey = "id"
const userPass = "test"
const userKey = "user"

type login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type User struct {
	Email string   `json:"email"`
	Roles []string `json:"roles"`
}

// AuthMiddleware creates a new JWT auth middleware
func AuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	key := make([]byte, hex.DecodedLen(len(APISecret)))
	_, err := hex.Decode(key, []byte(APISecret))
	if err != nil {
		return nil, err
	}

	mw, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "scw",
		Key:         key,
		Timeout:     30 * 24 * time.Hour,
		MaxRefresh:  30 * 24 * time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(User); ok {
				return jwt.MapClaims{
					identityKey: v.Email,
					"roles":     v.Roles,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)

			roles := make([]string, len(claims["roles"].([]interface{})))

			for i, role := range claims["roles"].([]interface{}) {
				roles[i] = role.(string)
			}
			email := claims[identityKey].(string)
			usr := User{
				Email: email,
				Roles: roles,
			}
			return usr
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login

			err := c.ShouldBind(&loginVals)
			if err != nil {
				log.Error(err)
				return "", jwt.ErrMissingLoginValues
			}

			usr := User{
				Email: loginVals.Email,
				Roles: []string{"admin", "user"},
			}

			if loginVals.Password != userPass {
				return nil, jwt.ErrFailedAuthentication
			}

			c.Set(userKey, usr)

			return usr, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if usr, ok := data.(User); ok && usr.Email == "harold@deckow.org" {
				return true
			}

			return false
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"code":   http.StatusOK,
				"token":  token,
				"user":   c.MustGet(userKey).(User),
				"expire": expire.Format(time.RFC3339),
			})
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	return mw, err
}
