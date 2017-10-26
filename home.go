package main

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/goadesign/goa"
	"github.com/jaredwarren/jlwarren1.com/app"
)

var templates = make(map[string]*template.Template)
var funcMap = template.FuncMap{
/*"title":  strings.Title,
"add":    add,
"tobool": tobool,*/
}

// Page ...
type Page struct {
	Title string
}

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

	templatePath := "home/index.html"
	// TODO: Move to outside or insice MakeMuxer func in production; user here to test, so templates are recompiled every request
	tpl := template.Must(template.New(templatePath).Funcs(funcMap).ParseFiles(fmt.Sprintf("templates/%s", templatePath), "templates/base.html"))

	var doc bytes.Buffer
	err := tpl.ExecuteTemplate(&doc, "base", nil)
	if err != nil {
		ctx.InternalServerError(err)
	}

	ctx.OK(doc.Bytes())

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

	templatePath := "home/resume.html"
	// TODO: Move to outside or insice MakeMuxer func in production; user here to test, so templates are recompiled every request
	tpl, err := template.New(templatePath).Funcs(funcMap).ParseFiles(fmt.Sprintf("templates/%s", templatePath), "templates/godocbase.html")
	tpl = template.Must(tpl, err)

	var doc bytes.Buffer
	page := &Page{
		Title: "Resume - Jared L. Warren",
	}
	err = tpl.ExecuteTemplate(&doc, "base", page)
	if err != nil {
		ctx.InternalServerError(err)
	}

	ctx.OK(doc.Bytes())

	// HomeController_Resume: end_implement
	return nil
}
