package topaz

import (
	"net/http"
)

// HTTP request handler function.
type Handler func(req Request, res Response)

type Server interface {
	// Creates a new handler for an endpoint accessed by a GET request. Path
	// can be formatted according to topaz url format standardds. Get the
	// params with the Request.PathParam() method.
	Get(path string, handler Handler)

	// Creates a new handler for an endpoint accessed by a POST request. Path
	// can be formatted according to topaz url format standards. Get the
	// params with the Request.PathParam method.
	Post(path string, handler Handler)

	// Serves a static site from the directory dir. Error is returned if
	// the directory does not exist.
	Static(path string, dir string) error

	// Serves the file directory dir to the given path. Error is returned if
	// the directory does not exist.
	ServeFiles(path string, dir string) error

	// Sets up server and listens to the port.
	Listen(port string) error

	// If the PORT environment variable is set it will return that, if not it
	// will fall back on the one given.
	EnvPort(port string) string
}

type Request interface {
	// Tries to unmarshal incoming data as json into dest.
	JSON(dest any) error

	// Returns the full URL as a string
	URL() string

	// Redirects the request to the new path.
	Redirect(path string)

	// Gets the value of a URL query parameter if present, otherwise an empty
	// string is returned.
	Query(key string) string

	// Gets the path parameters specified by the request handler path pattern.
	// Example path: "users/:id". Here you can get an "id" value from the url.
	// Returns an empty string if no param was found.
	Param(param string) string

	// Get the underlying http request object
	Request() *http.Request
}

type Response interface {
	// Writes bytes to response writer.
	Write(c []byte)

	// Responds to request with content as json data. Returns error if the
	// marshal failed, sends a server error status if so.
	JSON(content any) error

	// Responds to request with a given status. Status 200 is automatically
	// applied if the handler function does not fail. If another response function
	// has already succeeded, the new status will be ignored.
	Status(status int)

	// Responds to request with a file. Returns error if file does not exist.
	File(filename string) error

	// Get the underlying http response object
	Response() http.ResponseWriter
}
