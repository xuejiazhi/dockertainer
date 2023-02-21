package endpoint

import (
	"dockertainer/api/common"
	"dockertainer/api/databases"
	"dockertainer/api/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NodeList(c *gin.Context) {

}

func AddNode(c *gin.Context) {
	//get post value
	nodeName := c.DefaultPostForm("node_name", "localhost")
	nodeUrl := c.DefaultPostForm("node_url", "127.0.0.1:2379")
	nodeIp := c.DefaultPostForm("node_ip", "127.0.0.1")
	version := c.DefaultPostForm("ver", "v1.39")

	//判断传参
	if nodeName == "" || nodeUrl == "" || nodeIp == "" {
		c.JSON(http.StatusOK, common.NormalMsg{
			Code: http.StatusNoContent,
			Msg:  common.Tips["data_style_wrong"],
		})
		return
	}

	//校验IP
	if ok := util.RegexpIp(nodeIp); !ok {
		c.JSON(http.StatusOK, common.NormalMsg{
			Code: http.StatusNoContent,
			Msg:  common.Tips["wrong_ip_style"],
		})
		return
	}

	//增加节点
	if err := databases.AddNode(&databases.TNodePoint{
		NodeName: nodeName,
		NodeUrl:  nodeUrl,
		NodeIp:   nodeIp,
		Version:  version,
	}); err == nil {
		c.JSON(http.StatusOK, common.NormalMsg{
			Code: http.StatusOK,
			Msg:  common.Tips["add_node_succ"],
		})
	} else {
		c.JSON(http.StatusOK, common.NormalMsg{
			Code: http.StatusNoContent,
			Msg:  common.Tips["add_node_wrong"] + err.Error(),
		})
	}

}

func RemoveNode(c *gin.Context) {

}

func ModifyNode(c *gin.Context) {

}
