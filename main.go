package topaz

import "net/http"

// https://expressjs.com/en/4x/api.html

type Request interface {
	// Tries to unmarshal incoming data as json into obj.
	JSON(dest interface{}) error

	// Returns the full URL as a string
	URL() string

	// Gets the value of a URL query parameter if present, otherwise and error
	// is returned.
	Query(param string) (value string, err error)

	// Gets the path parameters specified by the request handler path pattern.
	// Example path: "users/:id". Here you can get an "id" value from the url.
	Param(param string) (value string, err error)

	// Get the underlying http request object
	Request() *http.Request
}

type Response interface {
	// Responds to request with content as json data. Returns error if the
	// marshal failed, sends a server error status if so.
	JSON(content interface{}) error

	// Responds to request with a given status. Status 200 is automatically
	// applied if the handler function does not fail.
	Status(status int)

	// Get the underlying http response object
	Response() http.ResponseWriter
}

// HTTP request handler function. Writes response data to response object.
type Handler func(req Request, res Response)

type Server interface {
	// Creates a new handler for an endpoint accessed by a GET request. Path
	// can be formatted according to pyrite url format standardds. Get the
	// params with the Request.PathParam() method.
	Get(path string, handlerFunc Handler)

	// Creates a new handler for an endpoint accessed by a POST request. Path
	// can be formatted according to pyrite url format standardds. Get the
	// params with the Request.PathParam() method.
	Post(path string, handlerFunc Handler)

	// Serves a static site from the entry point. Only serves files mentioned
	// by the entry point and its references.
	Static(entryPoint string) error

	// Sets up server and listens to the port. Canceled by either closing the
	// program or running Server.Close()
	Listen(port string) error

	// Stops server and gracefully ends all ongoing requests. Returns an error
	// if it has already been called.
	Close()
}
