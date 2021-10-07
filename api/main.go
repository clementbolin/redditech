package main

import (
	"github.com/redditech/api"
	"github.com/redditech/configs"
)

func main() {
	config := &configs.Config{}
	config.InitConfig("./configs/config.json")
	config.ReadConfig()
	router := api.InitRouter()
	router.Run(":8080")
}
