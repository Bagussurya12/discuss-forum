package main

import (
	"github.com/Bagussurya12/discuss-forum/source/handlers/memberships"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	membershipHandler := memberships.Newhandler(r)
	membershipHandler.RegisterRoute()
	r.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
