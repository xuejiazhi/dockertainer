package endpoint

import (
	"dockertainer/api/common"
	"dockertainer/api/databases"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetDashBoardInfo 获取面版详细信息
func GetDashBoardInfo(c *gin.Context) {
	//get params
	nodeName := c.DefaultQuery("node_name", "")

	//校验
	var nodeInfo databases.TNodePoint
	if msg, err := judgeNode(nodeName, &nodeInfo); err != nil {
		c.JSON(http.StatusOK, msg)
		return
	}

	//rest url get body
	restUrl := fmt.Sprintf("http://%s/v1.39/info", nodeInfo.NodeUrl)
	var dockerInfo DockerInfo
	if retData, err := getDockerApi(restUrl, &dockerInfo); err == nil {
		//返回
		c.JSON(http.StatusOK, retData)
	} else {
		c.JSON(http.StatusOK, common.NormalMsg{
			Code: http.StatusNoContent,
			Msg:  common.Tips["get_data_fail"],
		})
	}
}
