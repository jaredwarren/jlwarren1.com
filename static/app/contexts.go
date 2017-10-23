// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "jlwarren1": Application Contexts
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

// DashboardHomeContext provides the home dashboard action context.
type DashboardHomeContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewDashboardHomeContext parses the incoming request URL and body, performs validations and creates the
// context used by the home controller dashboard action.
func NewDashboardHomeContext(ctx context.Context, r *http.Request, service *goa.Service) (*DashboardHomeContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := DashboardHomeContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *DashboardHomeContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/html")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *DashboardHomeContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// ProfileHomeContext provides the home profile action context.
type ProfileHomeContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewProfileHomeContext parses the incoming request URL and body, performs validations and creates the
// context used by the home controller profile action.
func NewProfileHomeContext(ctx context.Context, r *http.Request, service *goa.Service) (*ProfileHomeContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ProfileHomeContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ProfileHomeContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/html")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ProfileHomeContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// ResumeHomeContext provides the home resume action context.
type ResumeHomeContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewResumeHomeContext parses the incoming request URL and body, performs validations and creates the
// context used by the home controller resume action.
func NewResumeHomeContext(ctx context.Context, r *http.Request, service *goa.Service) (*ResumeHomeContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ResumeHomeContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ResumeHomeContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/html")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ResumeHomeContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}
