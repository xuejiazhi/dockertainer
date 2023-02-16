package endpoint

import (
	"dockertainer/api/common"
	"dockertainer/api/databases"
	"dockertainer/api/util"
	"errors"
	"github.com/goccy/go-json"
	"net/http"
)

// ------------------------------------------------

type DockerMsgType interface {
	DockerInfo | ImageList | []ImageList | ImageInspect
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

func getDockerApi[T DockerMsgType](url string, valStruct *T) (msg common.ValueMsg, err error) {
	data := util.HttpGet(url)
	if err = json.Unmarshal([]byte(data), &valStruct); err == nil {
		msg.Code = http.StatusOK
		msg.Msg = common.Tips["get_data_succ"]
		msg.Value = valStruct
		return
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
