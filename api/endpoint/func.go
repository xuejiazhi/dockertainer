package endpoint

import (
	"dockertainer/api/common"
	"dockertainer/api/databases"
	"dockertainer/api/util"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
)

type NodeImageInfo struct {
	NodeName string
	ImageID  string
	NodeInfo databases.TNodePoint
	Msg      common.NormalMsg
}

// ------------------------------------------------

type DockerMsgType interface {
	DockerInfo | ImageList | []ImageList | ImageInspect | ImageHistory | []ImageHistory
}

func judgeNode(nodeName string, nodeInfo *databases.TNodePoint) (msg common.NormalMsg, err error) {
	if nodeName == "" {
		msg.Code = http.StatusNoContent
		msg.Msg = common.Tips["node_is_null"]
		err = errors.New(msg.Msg)
		return
	}

	//获取节点信息
	*nodeInfo = databases.GetNodeByName(nodeName)
	if nodeInfo.NodeUrl == "" {
		msg.Code = http.StatusNoContent
		msg.Msg = common.Tips["node_is_wrong"]
		err = errors.New(msg.Msg)
	}

	return
}

func judgeNodeId(c *gin.Context) (retData NodeImageInfo, err error) {
	//get params
	nodeName := c.DefaultQuery("node_name", "")
	retData.NodeName = nodeName
	//校验
	var nodeInfo databases.TNodePoint
	msg, err := judgeNode(nodeName, &nodeInfo)
	if err != nil {
		retData.Msg = msg
		return
	}

	//get params
	imageId := c.DefaultQuery("image_id", "")
	retData.ImageID = imageId
	if imageId == "" {
		msg = common.NormalMsg{
			Code: http.StatusNoContent,
			Msg:  common.Tips["id_is_null"],
		}
		err = errors.New(common.Tips["id_is_null"])
		retData.Msg = msg
	}

	retData.NodeInfo = nodeInfo

	return
}

func getDockerApi[T DockerMsgType](url string, valStruct *T) (msg common.ValueMsg, err error) {
	//获取Get respData
	data := util.HttpGet(url)
	if err = json.Unmarshal([]byte(data), &valStruct); err == nil {
		msg.Code = http.StatusOK
		msg.Msg = common.Tips["get_data_succ"]
		msg.Value = valStruct
	} else {
		msg.Code = http.StatusNoContent
		msg.Msg = common.Tips["get_data_fail"]
	}
	return
}

func deleteDockerApi(url string) (msg common.ValueMsg) {
	val := util.HttpDelete(url)
	msg.Code = http.StatusOK
	msg.Msg = common.Tips["operate_succ"]
	//值
	var tempMap map[string]interface{}
	json.Unmarshal([]byte(val), &tempMap)
	msg.Value = tempMap

	if val == "" {
		msg.Code = http.StatusNoContent
		msg.Msg = common.Tips["operate_fail"]
	}
	return
}
