package main

import (
	"log"

	"github.com/Bagussurya12/discuss-forum/source/configs"
	"github.com/Bagussurya12/discuss-forum/source/handlers/memberships"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./source/configs/"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("Failed Configuration:", err)
	}

	cfg = configs.Get()
	log.Println("Config", cfg)

	membershipHandler := memberships.Newhandler(r)
	membershipHandler.RegisterRoute()
	r.Run(cfg.Service.Port)
}
