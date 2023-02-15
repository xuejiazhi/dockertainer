package user

import (
	"dockertainer/api/common"
	"dockertainer/api/databases"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ModifyPassWord 更新用户密钥
func ModifyPassWord(c *gin.Context) {
	//get post value
	userName := c.PostForm("username")
	passWord := c.PostForm("password")

	if userName == "" {
		c.JSON(http.StatusOK, common.NormalMsg{
			Code: http.StatusNoContent,
			Msg:  "用户名称不能为空",
		})
		return
	}

	if len(passWord) < 6 || len(passWord) > 12 {
		c.JSON(http.StatusOK, common.NormalMsg{
			Code: http.StatusNoContent,
			Msg:  "密码大于6位 小于12位 ",
		})
		return
	}

	//query db
	userInfo := databases.QueryUser(userName)
	if userInfo.UserName != "" {
		if err := databases.ModifyPassword(userName, passWord); err != nil {
			c.JSON(http.StatusOK, common.NormalMsg{
				Code: http.StatusNoContent,
				Msg:  "Modify Password Error " + err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, common.NormalMsg{
				Code: http.StatusOK,
				Msg:  "Modify Password Success",
			})
		}
	} else {
		c.JSON(http.StatusOK, common.NormalMsg{
			Code: http.StatusNoContent,
			Msg:  "Modify Password Error ",
		})
	}
	return
}
