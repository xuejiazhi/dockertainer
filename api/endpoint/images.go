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
func RemoveImage(c *gin.Context) {
	//校验
	data, err := judgeNodeId(c)
	if err != nil {
		c.JSON(http.StatusOK, data.Msg)
		return
	}

	//remove url
	restUrl := fmt.Sprintf("http://%s/v1.39/images/%s", data.NodeInfo.NodeUrl, data.ImageID)
	retData := deleteDockerApi(restUrl)
	//返回
	c.JSON(http.StatusOK, retData)

}

// InspectImage 查看镜像详细信息
func InspectImage(c *gin.Context) {
	//校验
	data, err := judgeNodeId(c)
	if err != nil {
		c.JSON(http.StatusOK, data.Msg)
		return
	}

	//remove url
	restUrl := fmt.Sprintf("http://%s/v1.39/images/%s/json", data.NodeInfo.NodeUrl, data.ImageID)
	var imageInspect ImageInspect
	retData, err := getDockerApi(restUrl, &imageInspect)
	if err != nil {
		//todo: print logs
		fmt.Println("Image History->", err.Error())
	}

	c.JSON(http.StatusOK, retData)
}

// HistoryImage 返回images 历史记录
func HistoryImage(c *gin.Context) {
	//校验
	data, err := judgeNodeId(c)
	if err != nil {
		c.JSON(http.StatusOK, data.Msg)
		return
	}

	//remove url
	restUrl := fmt.Sprintf("http://%s/v1.39/images/%s/history", data.NodeInfo.NodeUrl, data.ImageID)
	var ImageHistorys []ImageHistory
	retData, err := getDockerApi(restUrl, &ImageHistorys)
	if err != nil {
		//todo: print logs
		fmt.Println("Image History->", err.Error())
	}
	//返回
	c.JSON(http.StatusOK, retData)
}
