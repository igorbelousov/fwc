package web

import (
	"net/http"

	"github.com/valyala/fasthttp"

	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

// Respond converts a Go value to JSON and sends it to the client.
func Respond(ctx *fasthttp.RequestCtx, data interface{}, statusCode int) error {

	// Set the status code for the request logger middleware.
	// If the context is missing this value, request the service
	// to be shutdown gracefully.

	// If there is nothing to marshal then set status code and return.
	if statusCode == fasthttp.StatusNoContent {
		ctx.Response.SetStatusCode(statusCode)
		return nil
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	// Convert the response value to JSON.
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Set the content type and headers once we know marshaling has succeeded.
	ctx.Response.Header.Set("Content-Type", "application/json")

	// Write the status code to the response.
	ctx.Response.SetStatusCode(statusCode)

	// Send the result back to the client.
	if _, err := ctx.Write(jsonData); err != nil {
		return err
	}

	return nil
}

// RespondError sends an error reponse back to the client.
func RespondError(ctx *fasthttp.RequestCtx, err error) error {

	// If the error was of the type *Error, the handler has
	// a specific status code and error to return.
	if webErr, ok := errors.Cause(err).(*Error); ok {
		er := ErrorResponse{
			Error:  webErr.Err.Error(),
			Fields: webErr.Fields,
		}
		if err := Respond(ctx, er, webErr.Status); err != nil {
			return err
		}
		return nil
	}

	// If not, the handler sent any arbitrary error value so use 500.
	er := ErrorResponse{
		Error: http.StatusText(http.StatusInternalServerError),
	}
	if err := Respond(ctx, er, fasthttp.StatusInternalServerError); err != nil {
		return err
	}

	return nil
}
