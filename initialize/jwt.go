package initialize

import (
	"time"

	"github.com/awoyai/gin-temp/global"
	"github.com/awoyai/gin-temp/model/common"
	"github.com/awoyai/gin-temp/repo/system"
	"github.com/awoyai/gin-temp/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func JwtInit() *utils.JWTMiddleware {
	timeout := time.Hour
	if global.CONFIG.JWT.TimeOut > 0 {
		timeout = time.Duration(global.CONFIG.JWT.TimeOut) * time.Second
	}

	var err error
	jwt, err := utils.New(&utils.JWTMiddleware{
		Key:            []byte(global.CONFIG.JWT.Secret),
		Timeout:        timeout,
		Authorizator:   Authorizator,
		TokenLookup:    "cookie:jwt,header:Authorization,query:token",
		TimeFunc:       time.Now,
		CookieHTTPOnly: true,
	})

	if err != nil {
		global.LOG.Fatal("JwtInit", zap.Error(err))
	}
	return jwt
}

func Authorizator(ctx *gin.Context) bool {
	claim, ok := ctx.Get(utils.JWT_PAYLOAD_KEY)
	if !ok {
		return false
	}
	username, ok := claim.(utils.MapClaims)[utils.TOKEN_KEY_NAME].(string)
	if !ok {
		return false
	}
	user, err := system.UserRepo.Info(map[string]any{"username": username})
	if err != nil {
		return false
	}
	if user.Enable == common.EnableTypeClose {
		return false
	}
	ctx.Set(common.LoginUserKey, username)
	return true
}
