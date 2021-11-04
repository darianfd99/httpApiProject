package courses

import (
	"net/http"

	mooc "github.com/darianfd99/httpApiProject/internal"
	"github.com/gin-gonic/gin"
)

type createRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

//CreateHandler returns an HTTP handler for courses creation
func CreateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		course := mooc.NewCourse(req.ID, req.Name, req.Duration)
		if err := Save(ctx, course); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return

		}

		ctx.Status(http.StatusCreated)
	}
}
