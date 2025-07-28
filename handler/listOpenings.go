package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reykalencar/gopportunities/schemas"
)

// @BasePath /api/v1

// @Summary List openings
// @Description show a list of job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
