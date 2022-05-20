package main

import (
	"fmt"
	"github.com/tedcy/fdfs_client"
)

func main() {
	//初始化客户端	--- 	配置文件
	client, err := fdfs_client.NewClientWithConfig("web/config/client.conf")
	if err != nil {
		fmt.Println(err)
		return
	}

	str, err := client.UploadByFilename("1.jpg")

	fmt.Println(str, err)

}
