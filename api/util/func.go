package util

import (
	"crypto/md5"
	"fmt"
	"github.com/idoubi/goz"
	"io"
	"net"
	"regexp"
)

func HashSaltMd5(password, salt string) string {
	h := md5.New()
	_, _ = io.WriteString(h, password+salt)
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

func HttpPost(url string, params map[string]interface{}) string {
	cli := goz.NewClient()
	var resp *goz.Response
	var err error
	if len(params) > 0 {
		resp, err = cli.Post(url, goz.Options{
			FormParams: params,
		})
	} else {
		resp, err = cli.Post(url)
	}

	if err != nil {
		return ""
	}
	body, _ := resp.GetBody()
	return body.String()
}

func HttpDelete(url string) string {
	cli := goz.NewClient()
	resp, err := cli.Delete(url)
	if err != nil {
		return ""
	}
	body, _ := resp.GetBody()
	return body.String()
}

func RegexpIp(ip string) bool {
	ipReg := `^(((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4})`
	r, _ := regexp.Compile(ipReg)
	match := r.MatchString(ip)
	if match {
		return true
	}
	return false
}
