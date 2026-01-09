package router

import (
	"ticket-system/internal/handler"
	"ticket-system/internal/middleware"
	"ticket-system/internal/model"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(t *handler.TicketHandler, u *handler.UserHandler) *gin.Engine {
	r := gin.Default()

	// public routes
	r.POST("/login", u.Login)
	r.POST("/register", u.Register)

	// protected routes
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	ticket := auth.Group("/tickets")
	{
		ticket.POST("/", t.Create)
		ticket.GET("/", t.List)
		ticket.GET("/:id", t.GetById)
		ticket.PUT("/:id/status", middleware.RequireRoles(model.RoleAgent, model.RoleAdmin), t.UpdateStatus)
		ticket.PUT("/:id/assign", middleware.RequireRoles(model.RoleAdmin), t.Assign)
	}
	return r
}
