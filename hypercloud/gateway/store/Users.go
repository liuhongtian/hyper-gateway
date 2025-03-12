package store

import (
	"hyper-gateway/hypercloud/gateway/data"
	"log"
)

var Users = make(map[string]*data.User)

func init() {
	log.Println("init users")
	defer log.Println("init users done")
	Users["liuhongtian"] = &data.User{
		Name:      "刘洪天",
		Gender:    "male",
		LoginName: "liuhongtian",
		Email:     "liuhongtian@163.com",
		CellPhone: "13800138000",
		Wechat:    "liuhongtian",
	}
	Users["zhangwan"] = &data.User{
		Name:      "张婉",
		Gender:    "female",
		LoginName: "zhanngwan",
		Email:     "zhanngwan@163.com",
		CellPhone: "13800138000",
		Wechat:    "zhangwan",
	}
}
