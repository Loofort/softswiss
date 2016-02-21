package main

import (
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/httpkit/middleware"

	"github.com/loofort/softswiss/restapi/operations"
	"github.com/loofort/softswiss/restapi/operations/command"
	"github.com/loofort/softswiss/restapi/operations/resource"
)

// This file is safe to edit. Once it exists it will not be overwritten

func configureAPI(api *operations.BankAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	api.JSONConsumer = httpkit.JSONConsumer()

	api.JSONProducer = httpkit.JSONProducer()

	api.ResourceAccountItemHandler = resource.AccountItemHandlerFunc(func(params resource.AccountItemParams) middleware.Responder {
		return middleware.NotImplemented("operation resource.AccountItem has not yet been implemented")
	})
	api.ResourceAccountListHandler = resource.AccountListHandlerFunc(func() middleware.Responder {
		return middleware.NotImplemented("operation resource.AccountList has not yet been implemented")
	})
	api.ResourceAccountRegistartionHandler = resource.AccountRegistartionHandlerFunc(func(params resource.AccountRegistartionParams) middleware.Responder {
		return middleware.NotImplemented("operation resource.AccountRegistartion has not yet been implemented")
	})
	api.CommandDepositHandler = command.DepositHandlerFunc(func(params command.DepositParams) middleware.Responder {
		return middleware.NotImplemented("operation command.Deposit has not yet been implemented")
	})
	api.CommandTransferHandler = command.TransferHandlerFunc(func(params command.TransferParams) middleware.Responder {
		return middleware.NotImplemented("operation command.Transfer has not yet been implemented")
	})
	api.CommandWithdrawHandler = command.WithdrawHandlerFunc(func(params command.WithdrawParams) middleware.Responder {
		return middleware.NotImplemented("operation command.Withdraw has not yet been implemented")
	})

	api.ServerShutdown = func() {}
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return NewUIHandler(handler)
}

type UIHandler struct {
	doc      string
	redirect http.Handler
	static   http.Handler
	next     http.Handler
}

func NewUIHandler(next http.Handler) http.Handler {
	doc := "/doc/"
	return UIHandler{
		doc:      doc,
		redirect: http.RedirectHandler(doc, http.StatusMovedPermanently),
		static:   http.StripPrefix(doc, http.FileServer(http.Dir("./swagger-ui/dist/"))),
		next:     next,
	}
}

func (h UIHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		h.redirect.ServeHTTP(rw, r)
		return
	}
	if r.URL.Path[:len(h.doc)] == h.doc {
		h.static.ServeHTTP(rw, r)
		return
	}
	h.next.ServeHTTP(rw, r)
}
