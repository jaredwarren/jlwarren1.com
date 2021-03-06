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

	Action("resume", func() {
		Description("Single Page Resume")
		Routing(GET("/resume"))
		Response(OK, "text/html")
		Response(InternalServerError, ErrorMedia)
	})

	Action("updateResume", func() {
		Description("Download resume pdf")
		Routing(GET("/updateresume"))
		//Response(OK, "text/html")
		Response(Created, "/resume")
		Response(InternalServerError, ErrorMedia)
	})

	Files("/resume/JaredWarren-Resume.pdf", "static/JaredWarren-Resume.pdf")
	Files("/images/favicon.ico", "static/images/favicon.ico")
	Files("/.well-known/keybase.txt", "static/keybase.txt")

	Files("/static/*filename", "static/")
})
