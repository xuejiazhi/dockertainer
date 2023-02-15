package endpoint

import (
	"dockertainer/api/common"
	"dockertainer/api/databases"
	"dockertainer/api/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
)

// GetDashBoardInfo 获取面版详细信息
func GetDashBoardInfo(c *gin.Context) {
	nodeName := c.DefaultQuery("node_name", "")
	if nodeName == "" {
		c.JSON(http.StatusOK, common.NormalMsg{
			Code: http.StatusNoContent,
			Msg:  "节点名称为空",
		})
		return
	}

	nodeInfo := databases.GetNodeByName(nodeName)
	if nodeInfo.NodeUrl == "" {
		c.JSON(http.StatusOK, common.NormalMsg{
			Code: http.StatusNoContent,
			Msg:  "节点信息错误",
		})
		return
	}

	//rest url
	restUrl := fmt.Sprintf("http://%s/v1.39/info", nodeInfo.NodeUrl)
	data := util.HttpGet(restUrl)
	var dockerInfo DockerInfo
	json.Unmarshal([]byte(data), &dockerInfo)
	c.JSON(http.StatusOK, dockerInfo)
}
