package main

import (
	"fmt"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
)



func main(){
	c := &config.AppConfig{
		AppID: "org-system-microservice-demoA",
		Cluster: "PRO",
		IP: "http://192.168.66.200:8080",
		NamespaceName: "application",
		IsBackupConfig: true,
		Secret: "b211e44ba7c54ab69ca4dc3e98578667",
	}

	client, _ := agollo.StartWithConfig(func() (*config.AppConfig, error){
		return c, nil
	})
	fmt.Println("初始化Apollo配置成功...")

	//用户开始获取apollo配置
	cache := client.GetConfigCache(c.NamespaceName)
	mysqlHost, _ := cache.Get("mysql.host")
	mysqlUsername, _ := cache.Get("mysql.username")
	mysqlPassword, _ := cache.Get("mysql.password")
	mysqlPort, _ := cache.Get("mysql.port")
	mysqlDb, _ := cache.Get("mysql.db")
	fmt.Println(mysqlHost)
	fmt.Println(mysqlUsername)
	fmt.Println(mysqlPassword)
	fmt.Println(mysqlPort)
	fmt.Println(mysqlDb)
}