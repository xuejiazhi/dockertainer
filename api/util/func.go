package util

import (
	"crypto/md5"
	"fmt"
	"github.com/idoubi/goz"
	"io"
	"net"
)

func HashSaltMd5(password, salt string) string {
	h := md5.New()
	io.WriteString(h, password+salt)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func GetMac() (mac string) {
	// 获取本机的MAC地址
	interfaces, err := net.Interfaces()
	if err != nil {
		return
	}
	for _, inter := range interfaces {
		if len(inter.HardwareAddr.String()) > 12 {
			mac = inter.HardwareAddr.String() //获取本机MAC地址
		}
	}
	return
}

func HttpGet(url string) string {
	cli := goz.NewClient()
	resp, err := cli.Get(url)
	if err != nil {
		return ""
	}
	body, _ := resp.GetBody()
	return body.String()
}
