package controllers

import (
	"cmdb/ent"
	"cmdb/ent/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

const TokenExpireDuration = time.Hour * 2
var JWTSecret = []byte("fdsaf43dsfd32432")

type MyClaims struct {
	UserName string `json:"user_name"`
	Role user.Role `json:"role"`
	jwt.StandardClaims
}

type LoginIn struct {
	Body ent.User
}

func Login(c *gin.Context, in *LoginIn)  {
}





