package main

import (
	"fmt"
	"net"
	"regexp"
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

func Test_regIP(t *testing.T) {
	str := "192.168.3.465666"
	ipReg := `^(((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4})`
	r, _ := regexp.Compile(ipReg)
	match := r.MatchString(str)
	if match {
		fmt.Printf("%s is a legal ipv4 address\n", str)
	} else {
		fmt.Printf("%s is not a legal ipv4 address\n", str)
	}
}
