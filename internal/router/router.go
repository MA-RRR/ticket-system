package router

import (
	"ticket-system/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(h *handler.TicketHandler) *gin.Engine {
	r := gin.Default()

	ticket := r.Group("/tickets")
	{
		ticket.POST("/", h.Create)
		ticket.GET("/", h.List)
		ticket.GET("/:id", h.GetById)
		ticket.PUT("/:id/status", h.UpdateStatus)
	}
	return r
}
