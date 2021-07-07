package util

import (
	"FunnyServer/global"
	"fmt"
	"strconv"
)

import "FunnyServer/consts"

import "golang.org/x/crypto/bcrypt"

//根据postAliasIndex和树洞号holeId为用户随机生成昵称
func GenerateAliasByReplyIndexAndHoleId(postAliasIndex uint, holeId int) (alias string) {
	//index := consts.HASH_HOLEID_FACTOR*hole_id + consts.HASH_REPLYINDEX_FACTOR*int(post_alias_index)
	name_length := len(consts.NAME)
	start_index := holeId * holeId % name_length
	//三种情况，依据昵称库进行生成
	//如果目前的postAliasIndex小于昵称总数，即当前发言的人均可被分配一个昵称，所以直接做一个简单的hash后产生新昵称
	//若postAliasIndex大于当前昵称数，小于昵称总数的平方，将两个昵称进行拼接
	//postAliasIndex大于昵称数的平方，则随即昵称后加hash字符的后五位
	if int(postAliasIndex) < name_length {
		alias_index := (start_index+int(postAliasIndex))%(name_length-1) + 1
		alias = consts.NAME[alias_index]
	} else if int(postAliasIndex)-name_length < (name_length-1)*(name_length-1) {
		//start_index % ((name_length - 1) * (name_length - 1))
		alias_index := (start_index+int(postAliasIndex))%((name_length-1)*(name_length-1)) + 1

		name1 := consts.NAME[(alias_index / (name_length - 1))]

		//from 1 to name_length-1
		name2 := consts.NAME[alias_index%(name_length-1)+1]
		alias = name1 + "与" + name2

	} else {
		alias_index := (start_index+int(postAliasIndex))%(name_length-1) + 1
		name := consts.NAME[alias_index]
		hashStr := HashWithSalt(strconv.Itoa(alias_index))
		hashSuffix := hashStr[len(hashStr)-5:]
		alias = name + hashSuffix
	}
	return
}


func HashWithSalt(plainText string) (HashText string) {

	hash, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.MinCost)
	if err!=nil{
		fmt.Println(err)
	}
	HashText = string(hash)
	return
}

func GenerateIntFromIp(ip string) (res uint) {
	res =0
	for i := 0; i < len(ip); i++ {
		if ip[i]>'0'&&ip[i]<'9'{
			t,_:=strconv.Atoi(string(ip[i]))
			res+=uint(t)
		}
	}
	return res% uint(len(consts.NAME))
}

func DecideMsgInsertIntoRedis(msg string)  {
	if len(msg)>len(consts.LEAVE_SUFFIX) && msg[len(msg)-len(consts.LEAVE_SUFFIX):]==consts.LEAVE_SUFFIX{
		return
	}
	if len(msg)>len(consts.COME_SUFFIX) && msg[len(msg)-len(consts.COME_SUFFIX):]==consts.COME_SUFFIX{
		return
	}
	count,err:=global.RedisClient.LLen("chat").Result()
	if err != nil {
		fmt.Println(err)
	}
	if count>100{
		go global.RedisClient.LPop("chat")
	}
	go global.RedisClient.RPush("chat",msg)
}

