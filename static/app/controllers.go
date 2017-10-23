// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "jlwarren1": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/jaredwarren/jlwarren1.com/design
// --force=true
// --out=$(GOPATH)/src/github.com/jaredwarren/jlwarren1.com/static
// --version=v1.2.0-dirty

package app

import (
	"context"
	"github.com/goadesign/goa"
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
	Resume(*ResumeHomeContext) error
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
		rctx, err := NewDashboardHomeContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Dashboard(rctx)
	}
	service.Mux.Handle("GET", "/", ctrl.MuxHandler("dashboard", h, nil))
	service.LogInfo("mount", "ctrl", "Home", "action", "Dashboard", "route", "GET /")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewProfileHomeContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Profile(rctx)
	}
	service.Mux.Handle("GET", "/profile", ctrl.MuxHandler("profile", h, nil))
	service.LogInfo("mount", "ctrl", "Home", "action", "Profile", "route", "GET /profile")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewResumeHomeContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Resume(rctx)
	}
	service.Mux.Handle("GET", "/resume", ctrl.MuxHandler("resume", h, nil))
	service.LogInfo("mount", "ctrl", "Home", "action", "Resume", "route", "GET /resume")

	h = ctrl.FileHandler("/static/*filename", "static/")
	service.Mux.Handle("GET", "/static/*filename", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Home", "files", "static/", "route", "GET /static/*filename")

	h = ctrl.FileHandler("/bower_components/*filename", "static/bower_components/")
	service.Mux.Handle("GET", "/bower_components/*filename", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Home", "files", "static/bower_components/", "route", "GET /bower_components/*filename")

	h = ctrl.FileHandler("/images/favicon.ico", "static/images/favicon.ico")
	service.Mux.Handle("GET", "/images/favicon.ico", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Home", "files", "static/images/favicon.ico", "route", "GET /images/favicon.ico")

	h = ctrl.FileHandler("/index.html", "static/index.html")
	service.Mux.Handle("GET", "/index.html", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Home", "files", "static/index.html", "route", "GET /index.html")

	h = ctrl.FileHandler("/index", "static/index.html")
	service.Mux.Handle("GET", "/index", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Home", "files", "static/index.html", "route", "GET /index")

	h = ctrl.FileHandler("/src/*filename", "static/src/")
	service.Mux.Handle("GET", "/src/*filename", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Home", "files", "static/src/", "route", "GET /src/*filename")

	h = ctrl.FileHandler("/static/", "static/index.html")
	service.Mux.Handle("GET", "/static/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Home", "files", "static/index.html", "route", "GET /static/")

	h = ctrl.FileHandler("/bower_components/", "static/bower_components/index.html")
	service.Mux.Handle("GET", "/bower_components/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Home", "files", "static/bower_components/index.html", "route", "GET /bower_components/")

	h = ctrl.FileHandler("/src/", "static/src/index.html")
	service.Mux.Handle("GET", "/src/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Home", "files", "static/src/index.html", "route", "GET /src/")
}