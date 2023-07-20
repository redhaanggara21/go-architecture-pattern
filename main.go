package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"gorm.io/gorm/logger"
	"red21.id/learn/bengkel/internal/component"
	"red21.id/learn/bengkel/internal/config"
	"red21.id/learn/bengkel/internal/module/customer"
	"red21.id/learn/bengkel/internal/module/history"
	"red21.id/learn/bengkel/internal/module/vehicle"
)

func main() {
	conf := config.Get()
	dbConnection := component.GetConnection(conf)

	customerRepository := customer.NewRepository(dbConnection)
	vehicleRepository := vehicle.NewRepository(dbConnection)
	historyRepository := history.NewRepository(dbConnection)

	customerService := customer.NewService(customerRepository)
	vehicleService := vehicle.NewService(vehicleRepository, historyRepository)
	historyService := history.NewService(historyRepositor)

	app := fiber.New()
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "[${locals:requestid}] ${ip} - ${method} - ${status} - ${path}]\n",
	}))
	customer.NewApi(app, customerService, vehicleService, historyService)

	_ = app.Listen(conf.Srv.Host + ":" + conf.Srv.Port)
}
