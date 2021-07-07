package broadcaster

import "FunnyServer/global"

func Broadcaster()  {
	for  {
		select {
		case msg:=<-global.Messages:
			go global.RedisClient.RPush("chat",msg)
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