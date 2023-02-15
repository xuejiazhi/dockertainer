package main

import (
	"dockertainer/api/endpoint"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/idoubi/goz"
	"log"
	"testing"
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
