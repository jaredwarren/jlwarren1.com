//************************************************************************//
// API "jlwarren1": Application Controllers
//
// Generated with goagen v0.2.dev, command line:
// $ goagen
// --design=github.com/jaredwarren/jlwarren1.com/design
// --out=$(GOPATH)\src\github.com\jaredwarren\jlwarren1.com
// --version=v0.2.dev
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// HomeController is the controller interface for the Home actions.
type HomeController interface {
	goa.Muxer
	goa.FileServer
	Dashboard(*DashboardHomeContext) error
	Profile(*ProfileHomeContext) error
	SetType(*SetTypeHomeContext) error
}

// MountHomeController "mounts" a Home resource controller on the given service.
func MountHomeController(service *goa.Service, ctrl HomeController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDashboardHomeContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Dashboard(rctx)
	}
	service.Mux.Handle("GET", "/", ctrl.MuxHandler("Dashboard", h, nil))
	service.LogInfo("mount", "ctrl", "Home", "action", "Dashboard", "route", "GET /")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewProfileHomeContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Profile(rctx)
	}
	service.Mux.Handle("GET", "/settype", ctrl.MuxHandler("Profile", h, nil))
	service.LogInfo("mount", "ctrl", "Home", "action", "Profile", "route", "GET /settype")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewSetTypeHomeContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.SetType(rctx)
	}
	service.Mux.Handle("GET", "/settype", ctrl.MuxHandler("SetType", h, nil))
	service.LogInfo("mount", "ctrl", "Home", "action", "SetType", "route", "GET /settype")

	h = ctrl.FileHandler("/static/*filename", "static/")
	service.Mux.Handle("GET", "/static/*filename", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Home", "files", "static/", "route", "GET /static/*filename")

	h = ctrl.FileHandler("/favicon.ico", "static/favicon.ico")
	service.Mux.Handle("GET", "/favicon.ico", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Home", "files", "static/favicon.ico", "route", "GET /favicon.ico")

	h = ctrl.FileHandler("/static/", "static\\index.html")
	service.Mux.Handle("GET", "/static/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Home", "files", "static\\index.html", "route", "GET /static/")
}
