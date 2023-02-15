package main

import (
	"dockertainer/api/endpoint"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/idoubi/goz"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"testing"
	"time"
)

func Test_curl(t *testing.T) {
	url := "http://10.161.30.207:11004/v1.39/info"
	cli := goz.NewClient()
	resp, err := cli.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, _ := resp.GetBody()
	fmt.Println(body.String())
	var dockerInfo endpoint.DockerInfo
	json.Unmarshal([]byte(body.String()), &dockerInfo)
	fmt.Println(dockerInfo)
}

func Test_telnet(t *testing.T) {
	for i := 1; i < 255; i++ {
		ip := fmt.Sprintf("10.110.14.%d", i)
		address := net.JoinHostPort(ip, "22")
		conn, err := net.DialTimeout("tcp", address, 3*time.Second)
		if err != nil {
			fmt.Println("Ip Address->", address, " telnet failed")
		} else {
			if conn != nil {
				fmt.Println("Ip Address->", address, " telnet success")
				sshclient(ip, 22)
				_ = conn.Close()
			} else {
				fmt.Println("Ip Address->", address, " telnet failed")
			}
		}

	}
}

func sshclient(sshHost string, sshPort int) {
	fmt.Println("-------Begin SSH-------")
	config := &ssh.ClientConfig{
		Timeout:         time.Second, //ssh 连接time out 时间一秒钟, 如果ssh验证错误 会在一秒内返回，尽量短些
		User:            "root",
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以, 但是不够安全
		Auth:            []ssh.AuthMethod{ssh.Password("Bingo@1993")},
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}
	addr := fmt.Sprintf("%s:%d", sshHost, sshPort)
	if sshClient, err := ssh.Dial("tcp", addr, config); err != nil {
		fmt.Println("-> ssh client failed,", err)
	} else {
		if sshClient != nil {
			fmt.Println("-> ssh client success")
			_ = sshClient.Close()
		} else {
			fmt.Println("-> ssh client failed")
		}
	}
	fmt.Println("-------SSH END-------")
}
