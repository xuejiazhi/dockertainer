package databases

import "gorm.io/gorm"

type TNodePoint struct {
	gorm.Model
	NodeName string `gorm:"column(node_name);size(64);unique"`
	NodeUrl  string `gorm:"column(node_url);size(264)"`
	NodeIp   string `gorm:"column(node_ip);size(264)"`
	Version  string `gorm:"column(node_ip);size(65)"`
}

func AddNode(nodeInfo *TNodePoint) error {
	return DBConn.Create(nodeInfo).Error
}

func GetNodeByName(nodeName string) (nodePoint TNodePoint) {
	DBConn.First(&nodePoint, "node_name=?", nodeName)
	return
}
