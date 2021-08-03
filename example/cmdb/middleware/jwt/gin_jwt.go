package jwt

import (
	"cmdb/ent"
	"cmdb/ent/user"
	"cmdb/public"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)
var identityKey = "id"


type loginValues struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type Authorizatorer interface {
	Authorizator(data interface{}, c *gin.Context) bool
}

func NewAuthMiddleware(authorizator Authorizatorer) (*jwt.GinJWTMiddleware, error) {
	realm := public.GetConf().Jwt.Realm
	key := public.GetConf().Jwt.SecretKey
	timeout, err := time.ParseDuration(public.GetConf().Jwt.Timeout)
	if err != nil {
		log.Fatalln(err)
	}
	maxRefresh, err := time.ParseDuration(public.GetConf().Jwt.MaxRefresh)
	if err != nil {
		log.Fatalln(err)
	}

	tokenHeadName := "Bearer"
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:      realm,
		Key:        []byte(key),
		Timeout:    timeout,
		MaxRefresh: maxRefresh,
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"code":   http.StatusOK,
				"token":  tokenHeadName + " " + token,
				"expire": expire.Format(time.RFC3339),
			})
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			login := loginValues{}
			if err := c.ShouldBind(&login); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			u, err := public.GetDB().User.Query().Where(user.EmailEQ(login.UserName), user.Password(login.Password)).Only(c)
			if err != nil {
				return nil, err
			}
			if u.ID == 0 {
				return nil, jwt.ErrMissingLoginValues
			}
			return u, err

		},
		Authorizator: authorizator.Authorizator,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*ent.User); ok {
				return jwt.MapClaims{"user_name": v.Name}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			role := claims["user_name"].(string)
			return role
		},
		IdentityKey: identityKey,
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:         tokenHeadName,
		TimeFunc:              time.Now,
	})
}
