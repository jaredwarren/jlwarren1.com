//go:generate goagen bootstrap -d github.com/jaredwarren/jlwarren1.com/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/jaredwarren/jlwarren1.com/app"
)

func main() {
	// Create service
	service := goa.New("jlwarren1")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "home" controller
	c := NewHomeController(service)
	app.MountHomeController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
