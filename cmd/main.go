package main

import (
	"log"

	"github.com/Bagussurya12/discuss-forum/pkg/internalsql"
	"github.com/Bagussurya12/discuss-forum/source/configs"
	"github.com/Bagussurya12/discuss-forum/source/handlers/memberships"
	"github.com/Bagussurya12/discuss-forum/source/handlers/posts"
	membershipRepo "github.com/Bagussurya12/discuss-forum/source/repository/memberships"
	postRepo "github.com/Bagussurya12/discuss-forum/source/repository/posts"
	membershipSvc "github.com/Bagussurya12/discuss-forum/source/service/memberships"
	postSvc "github.com/Bagussurya12/discuss-forum/source/service/posts"
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

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	membershipRepo := membershipRepo.NewRepository(db)
	postRepo := postRepo.NewRepository(db)
	membershipService := membershipSvc.NewService(cfg, membershipRepo)
	postService := postSvc.NewService(cfg, postRepo)

	membershipHandler := memberships.Newhandler(r, membershipService)
	postHandler := posts.Newhandler(r, postService)
	membershipHandler.RegisterRoute()
	postHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
