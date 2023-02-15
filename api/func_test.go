package main

import (
	"fmt"
	"net"
	"testing"
)

func Test_mac(t *testing.T) {
	// 获取本机的MAC地址
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Poor soul, here is what you got: " + err.Error())
	}
	for _, inter := range interfaces {
		fmt.Println(inter)
		fmt.Println(inter.Name)
		mac := inter.HardwareAddr //获取本机MAC地址
		fmt.Println("MAC ===== ", mac)
	}
}
