package main

import (
	"log"

	"ticket-system/internal/handler"
	"ticket-system/internal/model"
	"ticket-system/internal/repository"
	"ticket-system/internal/router"
	"ticket-system/internal/service"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/ticket?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&model.Ticket{}, &model.User{}, &model.AuditLog{})

	ticketRepo := repository.NewTicketRepository(db)
	ticketService := service.NewTicketService(ticketRepo)
	ticketHandler := handler.NewTicketHandler(ticketService)

	auditRepo := repository.NewAuditRepository(db)
	auditService := service.NewAuditService(auditRepo)
	ticketService.SetAuditService(auditService)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := router.SetUpRouter(ticketHandler, userHandler)
	r.Run(":8080")
}
