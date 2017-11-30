package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("jlwarren1", func() {
	Title("Jared L. Warren")
	Description("Me")
	Contact(func() {
		Name("Jared L. Warren")
		Email("jlwarren1.com")
		URL("http://www.jlwarren1.com")
	})
	Docs(func() {
		Description("BladeHQ Shipping Service")
		URL("http://wiki.bhq.local/index.php/BladeHQ_Shipping_Service_Application")
	})
	Host("localhost:8080")
	Scheme("http")

	ResponseTemplate(Created, func(pattern string) {
		Description("Resource created")
		Status(301)
		Headers(func() {
			Header("Location", String, "href to created resource", func() {
				Pattern(pattern)
			})
		})
	})
})
