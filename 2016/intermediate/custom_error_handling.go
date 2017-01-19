package main

import (
	"log"
	"net/http"
)

// START OMIT
func handleServerError(err error, request *Request, response *Response) {
	if routerError, ok := err.(AppError); ok { // HL

		if routerError.HttpStatusCode() == http.StatusInternalServerError {
			log.Println(err.Error())
		}

		response.WriteHeaderAndEntity(routerError.HttpStatusCode(), routerError)
		return
	}
	handleServerError(
		NewAppError( // HL
			ErrUnknown,
			err.Error(),
			http.StatusInternalServerError),
		request,
		response)
}

// END OMIT
