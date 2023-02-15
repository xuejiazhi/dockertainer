package endpoint

import (
	"dockertainer/api/common"
	"dockertainer/api/databases"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ImagesJson  镜像列表
func ImagesJson(c *gin.Context) {
	//get params
	nodeName := c.DefaultQuery("node_name", "")

	//校验
	var nodeInfo databases.TNodePoint
	if msg, err := judgeNode(nodeName, &nodeInfo); err != nil {
		c.JSON(http.StatusOK, msg)
		return
	}

	//rest url get body
	restUrl := fmt.Sprintf("http://%s/v1.39/images/json", nodeInfo.NodeUrl)
	var imageList []ImageList
	if retData, err := getDockerApi(restUrl, &imageList); err == nil {
		//返回
		c.JSON(http.StatusOK, retData)
	} else {
		c.JSON(http.StatusOK, common.NormalMsg{
			Code: http.StatusNoContent,
			Msg:  common.Tips["get_data_fail"],
		})
	}
}

// ImportFileTar 导入镜像包
func ImportFileTar() {

}

// ExportTarFile 导出镜像tar包
func ExportTarFile() {

}

// RemoveImage 删除镜像包,分为一般删除和强制删除
func RemoveImage() {

}
