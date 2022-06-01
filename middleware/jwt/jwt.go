package JWTHandler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"mini-douyin/models"
	"mini-douyin/utils"
	"net/http"
	"strings"
	"time"
)

// 自定义私钥, 不校验token的页面, token有效期
var (
	secret   = []byte("dddjm")
	noVerity = []string{
		"/feed/",          // 视频流接口API
		"/user/register/", // 用户注册接口API
		"/user/login/",    // 用户登录API
	}
	effectTime = 2 * time.Hour
)

// GenerateToken 生成Token
func GenerateToken(claims *models.UserClaims) string {
	claims.ExpiresAt = time.Now().Add(effectTime).Unix()
	sign, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	if err != nil {
		panic(err)
	}
	return sign
}

// JwtVerity 校验Token
func JwtVerity(ctx *gin.Context) {
	// 判断是否需要校验
	if IsContainArr(noVerity, ctx.Request.RequestURI) {
		return
	}
	// 获取Header中的token
	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.JSON(http.StatusOK, utils.GetReturnData(gin.H{"message": "Token Need"}, "ERROR"))
		return
	}
	valid := ParseToken(token)
	// 鉴定Token有效期
	switch valid.(type) {
	case string:
		ctx.JSON(http.StatusOK, utils.GetReturnData(gin.H{"message": valid}, "ERROR"))
		return
	default:
		// 设置上下文
		ctx.Set("user", ParseToken(token))
	}
}

// ParseToken 解析Token
func ParseToken(tokenString string) interface{} {
	token, err := jwt.ParseWithClaims(tokenString, &models.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return "token valid"
	}
	claims, ok := token.Claims.(*models.UserClaims)
	if !ok {
		return "token valid"
	}
	return claims
}

// Refresh 更新Token (预留)
func Refresh(tokenString string) string {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &models.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		panic(err)
	}
	claims, ok := token.Claims.(*models.UserClaims)
	if !ok {
		panic("token valid")
	}
	jwt.TimeFunc = time.Now
	claims.StandardClaims.ExpiresAt = time.Now().Add(2 * time.Hour).Unix()
	return GenerateToken(claims)
}

// IsContainArr 判断路由是否需要校验Token
func IsContainArr(noVerityAddr []string, c string) bool {
	for _, noAddr := range noVerityAddr {
		if strings.Contains(c, noAddr) {
			return true
		}
	}
	return false
}
