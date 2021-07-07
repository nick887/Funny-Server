package settings

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Settings struct {
	RedisSettings RedisSetting `json:"RedisSettings"`
}

type RedisSetting struct {
	Address  string `json:"Address"`
	Password string `json:"Password"`
	Port     string `json:"Port"`
}

func ReadSettingsFromFile(settingFilePath string) (res *Settings) {
	var settings Settings
	jsonFile, err := os.Open(settingFilePath)
	if err != nil {
		panic("No such file named " + settingFilePath)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &settings)
	if err != nil {
		log.Panic(err)
	}
	res = &settings
	return
}