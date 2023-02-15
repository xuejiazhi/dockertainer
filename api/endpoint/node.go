package endpoint

import (
	"dockertainer/api/common"
	"dockertainer/api/databases"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NodeList(c *gin.Context) {

}

func AddNode(c *gin.Context) {
	//get post value
	nodeName := c.PostForm("node_name")
	nodeUrl := c.PostForm("node_url")
	nodeIp := c.PostForm("node_ip")

	//todo : 校验IP
	if nodeName == "" || nodeUrl == "" || nodeIp == "" {
		c.JSON(http.StatusOK, common.NormalMsg{
			Code: http.StatusNoContent,
			Msg:  "数据格式不正确",
		})
		return
	}

	//增加节点
	if err := databases.AddNode(&databases.TNodePoint{
		NodeName: nodeName,
		NodeUrl:  nodeUrl,
		NodeIp:   nodeIp,
	}); err == nil {
		c.JSON(http.StatusOK, common.NormalMsg{
			Code: http.StatusOK,
			Msg:  "增加节点成功",
		})
	} else {
		c.JSON(http.StatusOK, common.NormalMsg{
			Code: http.StatusNoContent,
			Msg:  "增加节点失败," + err.Error(),
		})
	}

}

func RemoveNode(c *gin.Context) {

}

func ModifyNode(c *gin.Context) {

}
