package restful

import "net/http"

type Request struct {
	Request        *http.Request // HL
	pathParameters map[string]string
	attributes     map[string]interface{} // for storing request-scoped values
}
