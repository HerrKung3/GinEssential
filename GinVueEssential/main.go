package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"herrkung.com/GinVueEssential/common"
)

func main() {
	InitConfig()
	common.InitDB()
	r := gin.Default()
	r = CollectRoutes(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
