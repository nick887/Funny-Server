package global

import (
	"FunnyServer/settings"
	"github.com/go-redis/redis"
)


var RedisClient *redis.Client


var (
	Settings *settings.Settings
)

type Client chan<- string

var (
	Entering = make(chan Client)
	Leaving = make(chan Client)
	Messages = make(chan string)
	Clients = make(map[Client]bool)
)