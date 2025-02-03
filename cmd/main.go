package main

import (
	"log"

	"github.com/Bagussurya12/discuss-forum/pkg/internalsql"
	"github.com/Bagussurya12/discuss-forum/source/configs"
	"github.com/Bagussurya12/discuss-forum/source/handlers/memberships"
	membershipRepo "github.com/Bagussurya12/discuss-forum/source/repository/memberships"
	membershipSvc "github.com/Bagussurya12/discuss-forum/source/service/memberships"
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

	db, err := internalsql.Connect(cfg.Database.DataSourceName)

	if err != nil {
		log.Fatal("Failed Initialitation Database: ", err)
	}

	membershipRepo := membershipRepo.NewRepository(db)
	membershipService := membershipSvc.NewService(membershipRepo)

	membershipHandler := memberships.Newhandler(r, membershipService)
	membershipHandler.RegisterRoute()
	r.Run(cfg.Service.Port)
}
