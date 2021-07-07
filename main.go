package main

import (
	"FunnyServer/broadcaster"
	"FunnyServer/global"
	"FunnyServer/handler"
	"FunnyServer/settings"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"net"
)

func main() {
	listener,err:=net.Listen("tcp",":9999")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster.Broadcaster()
	for  {
		conn,err:=listener.Accept()
		fmt.Println(conn.RemoteAddr().String()+"  connected")
		go global.RedisClient.LPush("log",conn.RemoteAddr().String()+"  connected")
		if err!=nil{
			log.Print(err)
			continue
		}
		go handler.HandleConn(conn)
	}
}



func init() {
	initSetting()
	initRedis()
}

func initRedis()  {
	redisSetting := global.Settings.RedisSettings
	global.RedisClient = redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:      redisSetting.Address + ":" + redisSetting.Port,
		Password: redisSetting.Password,
		DB:       0,
	})
	pong, err := global.RedisClient.Ping().Result()
	fmt.Println(pong, err)
}

func initSetting()  {
	global.Settings=settings.ReadSettingsFromFile("Settings.json")
}