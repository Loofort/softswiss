package main

import (
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/httpkit/middleware"

	"github.com/loofort/softswiss/models"
	"github.com/loofort/softswiss/restapi/operations"
	"github.com/loofort/softswiss/restapi/operations/command"
	"github.com/loofort/softswiss/restapi/operations/resource"
	"github.com/loofort/softswiss/storage"
)

// This file is safe to edit. Once it exists it will not be overwritten

func configureAPI(api *operations.BankAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	api.JSONConsumer = httpkit.JSONConsumer()

	api.JSONProducer = httpkit.JSONProducer()

	stg := storage.MustConnect("")

	/********************** Recource Handlers **************************/
	api.ResourceAccountItemHandler = resource.AccountItemHandlerFunc(func(params resource.AccountItemParams) middleware.Responder {
		account, err := stg.AccountItem(params.ID)
		if storage.IsNotFound(err) {
			return resource.NewAccountItemNotFound().WithPayload(&models.Error{err.Error()})
		}
		if err != nil {
			return resource.NewAccountItemDefault(0).WithPayload(&models.Error{err.Error()})
		}
		return resource.NewAccountItemOK().WithPayload(account)
	})

	api.ResourceAccountListHandler = resource.AccountListHandlerFunc(func() middleware.Responder {
		accounts, err := stg.AccountList()
		if err != nil {
			return resource.NewAccountListDefault(0).WithPayload(&models.Error{err.Error()})
		}
		return resource.NewAccountListOK().WithPayload(accounts)
	})

	api.ResourceAccountRegistartionHandler = resource.AccountRegistartionHandlerFunc(func(params resource.AccountRegistartionParams) middleware.Responder {
		acc := &models.Account{Amount: params.Body.Amount}
		acc, err := stg.AccountInsert(acc)
		if err != nil {
			return resource.NewAccountRegistartionDefault(0).WithPayload(&models.Error{err.Error()})
		}

		return resource.NewAccountRegistartionCreated().WithPayload(acc)
	})

	/********************** Command Handlers **************************/
	api.CommandDepositHandler = command.DepositHandlerFunc(func(params command.DepositParams) middleware.Responder {
		tx := stg.Begin()
		responder := cmdDepositHandler(params, tx)
		if _, ok := responder.(*command.DepositOK); !ok {
			tx.Rollback()
			return responder
		}
		tx.Commit()
		return responder
	})

	api.CommandWithdrawHandler = command.WithdrawHandlerFunc(func(params command.WithdrawParams) middleware.Responder {
		tx := stg.Begin()
		responder := cmdWithdrawHandler(params, tx)
		if _, ok := responder.(*command.WithdrawOK); !ok {
			tx.Rollback()
			return responder
		}
		tx.Commit()
		return responder
	})

	api.CommandTransferHandler = command.TransferHandlerFunc(func(params command.TransferParams) middleware.Responder {
		tx := stg.Begin()
		responder := cmdTransferHandler(params, tx)
		if _, ok := responder.(*command.TransferOK); !ok {
			tx.Rollback()
			return responder
		}
		tx.Commit()
		return responder
	})

	api.ServerShutdown = func() {}
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

func cmdDepositHandler(params command.DepositParams, stg storage.Tx) middleware.Responder {
	account, err := stg.AccountItem(params.Body.ID)
	if storage.IsNotFound(err) {
		return command.NewDepositNotFound().WithPayload(&models.Error{err.Error()})
	}
	if err != nil {
		return command.NewDepositDefault(0).WithPayload(&models.Error{err.Error()})
	}

	account.Amount += params.Body.Amount
	if err := stg.AccountUpdate(account); err != nil {
		return command.NewDepositDefault(0).WithPayload(&models.Error{err.Error()})
	}

	return command.NewDepositOK().WithPayload(account)
}

func cmdWithdrawHandler(params command.WithdrawParams, stg storage.Tx) middleware.Responder {
	account, err := stg.AccountItem(params.Body.ID)
	if storage.IsNotFound(err) {
		return command.NewWithdrawNotFound().WithPayload(&models.Error{err.Error()})
	}
	if err != nil {
		return command.NewWithdrawDefault(0).WithPayload(&models.Error{err.Error()})
	}

	account.Amount -= params.Body.Amount
	if account.Amount <= 0 {
		return command.NewWithdrawDefault(http.StatusNotAcceptable).WithPayload(&models.Error{"not enough money"})
	}

	if err := stg.AccountUpdate(account); err != nil {
		return command.NewWithdrawDefault(0).WithPayload(&models.Error{err.Error()})
	}

	return command.NewWithdrawOK().WithPayload(account)
}

func cmdTransferHandler(params command.TransferParams, stg storage.Tx) middleware.Responder {
	from, err := stg.AccountItem(params.Body.From)
	if storage.IsNotFound(err) {
		return command.NewTransferNotFound().WithPayload(&models.Error{err.Error()})
	}
	if err != nil {
		return command.NewTransferDefault(0).WithPayload(&models.Error{err.Error()})
	}

	to, err := stg.AccountItem(params.Body.To)
	if storage.IsNotFound(err) {
		return command.NewTransferNotFound().WithPayload(&models.Error{err.Error()})
	}
	if err != nil {
		return command.NewTransferDefault(0).WithPayload(&models.Error{err.Error()})
	}

	from.Amount -= params.Body.Amount
	if from.Amount <= 0 {
		return command.NewTransferDefault(http.StatusNotAcceptable).WithPayload(&models.Error{"not enough money"})
	}

	if err := stg.AccountUpdate(from); err != nil {
		return command.NewTransferDefault(0).WithPayload(&models.Error{err.Error()})
	}

	to.Amount += params.Body.Amount
	if err := stg.AccountUpdate(to); err != nil {
		return command.NewTransferDefault(0).WithPayload(&models.Error{err.Error()})
	}

	accs := []*models.Account{from, to}
	return command.NewTransferOK().WithPayload(accs)
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
