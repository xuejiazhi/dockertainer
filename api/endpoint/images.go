package endpoint

import (
	"dockertainer/api/common"
	"dockertainer/api/databases"
	"dockertainer/api/util"
	"encoding/json"
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

	//过滤器
	reference := c.DefaultQuery("reference", "")
	if reference != "" {
		var imageList []ImageList
		filters := fmt.Sprintf("*%s*:*", reference)
		//因为filters 不支持/,所以采用这笨办法
		for i := 0; i < 5; i++ {
			if i > 0 {
				filters = "*/" + filters
			}
			r, _ := json.Marshal(map[string][]string{
				"reference": {filters},
			})
			queryUrl := fmt.Sprintf("%s?filters=%s", restUrl, string(r))
			//获取GET数据
			var images []ImageList
			data := util.HttpGet(queryUrl)
			_ = json.Unmarshal([]byte(data), &images)
			imageList = append(imageList, images...)
		}
		//返回
		c.JSON(http.StatusOK, common.ValueMsg{
			Code:  http.StatusOK,
			Msg:   common.Tips["get_data_succ"],
			Value: imageList,
		})
	} else {
		//非过滤的情况
		var imageList []ImageList
		retData, err := getDockerApi(restUrl, &imageList)
		if err != nil {
			//todo: print logs
			fmt.Println("Image History->", err.Error())
		}
		c.JSON(http.StatusOK, retData)
	}
}

// ImagesSearch 镜像查询
func ImagesSearch(c *gin.Context) {
	//get params
	nodeName := c.DefaultQuery("node_name", "")

	//校验
	var nodeInfo databases.TNodePoint
	if msg, err := judgeNode(nodeName, &nodeInfo); err != nil {
		c.JSON(http.StatusOK, msg)
		return
	}

}

// ImagesTag  给镜像打tag
func ImagesTag(c *gin.Context) {
	//校验
	data, err := judgeNodeId(c)
	if err != nil {
		c.JSON(http.StatusOK, data.Msg)
		return
	}
	//获取repo tag信息
	repo := c.DefaultQuery("repo", "")
	if repo == "" {
		c.JSON(http.StatusOK, common.NormalMsg{
			Code: http.StatusNoContent,
			Msg:  common.Tips["repo_is_null"],
		})
		return
	}

	tag := c.DefaultQuery("tag", "latest")

	//remove url
	restUrl := fmt.Sprintf("http://%s/%s/images/%s/tag", data.NodeInfo.NodeUrl, data.NodeInfo.Version, data.ImageID)
	var retMap map[string]interface{}
	retData, err := postDockerApi(restUrl,
		map[string]interface{}{
			"Content-Type": "application/x-www-form-urlencoded",
		},
		map[string]interface{}{
			"repo": repo,
			"tag":  tag,
		},
		&retMap)
	if err != nil {
		//todo: print logs
		fmt.Println("Image Tag->", err.Error())
	}
	//返回
	c.JSON(http.StatusOK, retData)
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
	restUrl := fmt.Sprintf("http://%s/%s/images/%s", data.NodeInfo.NodeUrl, data.NodeInfo.Version, data.ImageID)
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
	restUrl := fmt.Sprintf("http://%s/%s/images/%s/json", data.NodeInfo.NodeUrl, data.NodeInfo.Version, data.ImageID)
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
	restUrl := fmt.Sprintf("http://%s/%s/images/%s/history", data.NodeInfo.NodeUrl, data.NodeInfo.Version, data.ImageID)
	var ImageHistorys []ImageHistory
	retData, err := getDockerApi(restUrl, &ImageHistorys)
	if err != nil {
		//todo: print logs
		fmt.Println("Image History->", err.Error())
	}
	//返回
	c.JSON(http.StatusOK, retData)
}
