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

	Action("profile", func() {
		Description("Set Printer Type")
		Routing(GET("/profile"))
		Response(OK, "text/html")
		Response(InternalServerError, ErrorMedia)
	})

	Files("/index.html", "static/index.html")
	Files("/index", "static/index.html")
	Files("/images/favicon.ico", "static/images/favicon.ico")
	Files("/static/*filename", "static/")
	Files("/src/*filename", "static/src/")
	Files("/bower_components/*filename", "static/bower_components/")

})
