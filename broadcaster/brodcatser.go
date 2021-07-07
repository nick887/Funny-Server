package broadcaster

import (
	"FunnyServer/global"
	"FunnyServer/util"
)

func Broadcaster()  {
	for  {
		select {
		case msg:=<-global.Messages:
			util.DecideMsgInsertIntoRedis(msg)
			for cli:=range global.Clients{
				cli<-msg
			}
		case cli:=<-global.Entering:
			global.Clients[cli] =true
		case cli:=<-global.Leaving:
			delete(global.Clients,cli)
			close(cli)
		}
	}
}
