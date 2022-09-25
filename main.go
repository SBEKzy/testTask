package main

import (
	"github.com/SBEKzy/testTask/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"

	"github.com/SBEKzy/testTask/handler"
	"github.com/SBEKzy/testTask/repository"
	"github.com/SBEKzy/testTask/service"
)

func main() {

	conf := config.Load()
	repo := repository.New(conf)
	s := service.New(repo)
	h := handler.New(s)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
	r := gin.Default()
	r.POST("/send", h.SendRequestForThirdService)
	err = r.Run(conf.Address)
	if err != nil {
		log.Fatalln(err)
	}

}
