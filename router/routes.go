package router

import (
	"net/http"
)

type Route struct {
	name string
	method string
	pattern string
	handlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes {
	Route {"Index", "GET", "/", Index},
	Route {"Index", "GET", "/accounts_rank/{number}", AccountsRank},
	Route {"Index", "GET", "/token_accounts_rank/{tokenName}", TokenAccountsRank},
	Route {"Index", "GET", "/api/token_accounts_rank/{tokenName}", ApiTokenAccountsRank},
	Route {"Index", "GET", "/api/accounts_rank/{number}", ApiAccountsRank},
}