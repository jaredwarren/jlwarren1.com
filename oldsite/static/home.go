package main

import (
	"github.com/goadesign/goa"
	"github.com/jaredwarren/jlwarren1.com/static/app"
)

// HomeController implements the home resource.
type HomeController struct {
	*goa.Controller
}

// NewHomeController creates a home controller.
func NewHomeController(service *goa.Service) *HomeController {
	return &HomeController{Controller: service.NewController("HomeController")}
}

// Dashboard runs the dashboard action.
func (c *HomeController) Dashboard(ctx *app.DashboardHomeContext) error {
	// HomeController_Dashboard: start_implement

	// Put your logic here

	// HomeController_Dashboard: end_implement
	return nil
}

// Profile runs the profile action.
func (c *HomeController) Profile(ctx *app.ProfileHomeContext) error {
	// HomeController_Profile: start_implement

	// Put your logic here

	// HomeController_Profile: end_implement
	return nil
}

// Resume runs the resume action.
func (c *HomeController) Resume(ctx *app.ResumeHomeContext) error {
	// HomeController_Resume: start_implement

	// Put your logic here

	// HomeController_Resume: end_implement
	return nil
}
