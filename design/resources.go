package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Home API
var _ = Resource("home", func() {
	BasePath("/")

	Action("dashboard", func() {
		Description("Show basic actions")
		Routing(GET(""))
		Response(OK, "text/html")
		Response(InternalServerError, ErrorMedia)
	})

	Action("setType", func() {
		Description("Set Printer Type")
		Routing(GET("/settype"))
		Response(OK, "text/html")
		Response(InternalServerError, ErrorMedia)
	})

	Action("profile", func() {
		Description("Set Printer Type")
		Routing(GET("/settype"))
		Response(OK, "text/html")
		Response(InternalServerError, ErrorMedia)
	})

	Files("/favicon.ico", "static/favicon.ico")
	Files("/static/*filename", "static/")

})
