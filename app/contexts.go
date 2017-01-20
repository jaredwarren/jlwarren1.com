//************************************************************************//
// API "jlwarren1": Application Contexts
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
)

// DashboardHomeContext provides the home dashboard action context.
type DashboardHomeContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewDashboardHomeContext parses the incoming request URL and body, performs validations and creates the
// context used by the home controller dashboard action.
func NewDashboardHomeContext(ctx context.Context, service *goa.Service) (*DashboardHomeContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
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
func NewProfileHomeContext(ctx context.Context, service *goa.Service) (*ProfileHomeContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
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

// SetTypeHomeContext provides the home setType action context.
type SetTypeHomeContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewSetTypeHomeContext parses the incoming request URL and body, performs validations and creates the
// context used by the home controller setType action.
func NewSetTypeHomeContext(ctx context.Context, service *goa.Service) (*SetTypeHomeContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := SetTypeHomeContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *SetTypeHomeContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/html")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *SetTypeHomeContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}
