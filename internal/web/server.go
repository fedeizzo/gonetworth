package web

import (
	"net/http"

	"github.com/oapi-codegen/runtime/strictmiddleware/nethttp"

	"github.com/fedeizzo/gonetworth/internal/web/api"
	"github.com/fedeizzo/gonetworth/internal/web/handler"
)

func Run() error {

	router := http.NewServeMux()

	// validation
	// swagger, err := api.GetSwagger()
	// if err != nil {
	// 	return nil, err
	// }

	// validator := middleware.OapiRequestValidator(swagger)
	// middleware := api.StrictMiddlewareFunc(func(f nethttp.StrictHTTPHandlerFunc, operationID string) nethttp.StrictHTTPHandlerFunc {

	// })

	handler := api.NewStrictHandler(&handler.Accounts{}, []nethttp.StrictHTTPMiddlewareFunc{})

	api.HandlerWithOptions(handler, api.StdHTTPServerOptions{
		BaseRouter: router,
	})

	return http.ListenAndServe("localhost:3000", router)
}
