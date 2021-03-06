package resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit/middleware"
)

// AccountItemHandlerFunc turns a function with the right signature into a account item handler
type AccountItemHandlerFunc func(AccountItemParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AccountItemHandlerFunc) Handle(params AccountItemParams) middleware.Responder {
	return fn(params)
}

// AccountItemHandler interface for that can handle valid account item params
type AccountItemHandler interface {
	Handle(AccountItemParams) middleware.Responder
}

// NewAccountItem creates a new http.Handler for the account item operation
func NewAccountItem(ctx *middleware.Context, handler AccountItemHandler) *AccountItem {
	return &AccountItem{Context: ctx, Handler: handler}
}

/*AccountItem swagger:route GET /account/{id} resource accountItem

get info about particular account

*/
type AccountItem struct {
	Context *middleware.Context
	Handler AccountItemHandler
}

func (o *AccountItem) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewAccountItemParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
