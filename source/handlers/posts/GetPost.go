package posts

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetPostById(c *gin.Context) {
	ctx := c.Request.Context()

	postIdStr := c.Param("postID")
	postId, err := strconv.ParseInt(postIdStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("invalid post id").Error(),
		})
		return
	}

	response, err := h.postSvc.GetPostById(ctx, postId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
