package user

import (
	"dockertainer/api/cache"
	"dockertainer/api/common"
	"dockertainer/api/databases"
	"dockertainer/api/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Auth 校验密钥的正确性
func Auth(c *gin.Context) {
	//get post value
	userName := c.PostForm("username")
	passWord := c.PostForm("password")

	//query db
	userInfo := databases.QueryUser(userName)
	if userInfo.UserName != "" && userInfo.PassWord == passWord {
		//设置缓存
		mac := util.GetMac()
		cache.Set(mac, generalRegValue(mac, userName, passWord),
			common.CacheKeyList["user_reg_key"].TTL)
		//返回
		c.JSON(http.StatusOK, common.NormalMsg{
			Code: http.StatusOK,
			Msg:  "Auth Success",
		})
	} else {
		c.JSON(http.StatusOK, common.NormalMsg{
			Code: http.StatusNoContent,
			Msg:  "Auth Error",
		})
	}
}

func generalRegValue(mac, user, password string) map[string]string {
	//获取mac地址
	retData := map[string]string{
		"user": user,
		"hex":  util.HashSaltMd5(password, mac),
	}
	return retData
}
