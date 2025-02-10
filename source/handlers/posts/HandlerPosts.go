package posts

import (
	"context"

	"github.com/Bagussurya12/discuss-forum/source/middleware"
	"github.com/Bagussurya12/discuss-forum/source/model/posts"
	"github.com/gin-gonic/gin"
)

type postService interface {
	CreatePost(ctx context.Context, userID int64, req posts.CreatePostRequest) error
	CreateComment(ctx context.Context, postID, userID int64, request posts.CreateCommentRequest) error
}

type Handler struct {
	*gin.Engine

	postSvc postService
}

func Newhandler(api *gin.Engine, postSvc postService) *Handler {
	return &Handler{
		Engine:  api,
		postSvc: postSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("posts")
	route.Use(middleware.AuthMiddeware())

	route.POST("/create-post", h.CreatePost)
	route.POST("/create-comment", h.CreateComment)
}
