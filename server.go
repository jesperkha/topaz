package topaz

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var (
	errNoDirectory = errors.New("could not find file or directory '%s'")
)

// Todo: new instance of server mux for each topaz server
func NewServer() Server {
	return &server{}
}

func (s *server) Get(path string, handlerFunc Handler) {
	s.handle(http.MethodGet, path, handlerFunc)
}

func (s *server) Post(path string, handlerFunc Handler) {
	s.handle(http.MethodPost, path, handlerFunc)
}

func (s *server) Static(path string, dir string) error {
	if _, err := os.Open(path); err != nil {
		return errorf(errNoDirectory, dir)
	}
	http.Handle(path, http.FileServer(http.Dir(dir)))
	return nil
}

func (s *server) Listen(port string) error {
	return http.ListenAndServe(port, nil)
}

type server struct {
}

type urlParams struct {
	parameters []parameter

	// Identifier path. Replaces the variable URL params with '-'
	// to check for matching paths of slightly different URLs.
	identifier string
}

// URL parameter object. Index is the position in the URL (split
// by '/') starting at 0 after the root.
type parameter struct {
	name  string
	index int
}

// Returns a split version of the raw URL path without a leading '/'
func pathSplit(url string) []string {
	return strings.Split(strings.Split(url, "?")[0], "/")[1:]
}

// Returns slice of paramter names in order
func getUrlParams(url string) urlParams {
	split := pathSplit(url)
	params := urlParams{}

	identifier := ""
	for idx, c := range split {
		if strings.HasPrefix(c, ":") {
			param := parameter{
				name:  c[1:],
				index: idx,
			}

			identifier += "/-"
			params.parameters = append(params.parameters, param)
		} else {
			identifier += "/" + c
		}
	}

	params.identifier = identifier
	return params
}

func resreq(w http.ResponseWriter, r *http.Request) (res *response, req *request) {
	return &response{response: w, request: r},
		&request{
			request:  r,
			response: w,
			params:   map[string]string{},
		}
}

func (s *server) handle(method string, path string, handlerFunc Handler) {
	params := getUrlParams(path)

	httpHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		res, req := resreq(w, r)
		urlSplit := pathSplit(r.URL.Path)
		for _, param := range params.parameters {
			// Index cannot be out of range if identifiers match in length
			req.params[param.name] = urlSplit[param.index]
			urlSplit[param.index] = "-"
		}

		// Non-matching url identifier patterns
		if fmt.Sprintf("/%s", strings.Join(urlSplit, "/")) != params.identifier {
			return
		}

		handlerFunc(req, res)
		if !req.redirected && res.status == 0 {
			w.WriteHeader(http.StatusOK)
		}
	}

	// Get leading path before any variable url parameters
	rootPath := "/"
	if split := pathSplit(path); len(split) > 1 {
		rootPath = ""
		for _, s := range split {
			if !strings.HasPrefix(s, ":") {
				rootPath += "/" + s
			} else {
				rootPath += "/"
				break
			}
		}
	}

	http.HandleFunc(rootPath, httpHandler)
}

func (s *server) ServeFiles(path string, dir string) {
	http.Handle(path, http.StripPrefix(path, http.FileServer(http.Dir(dir))))
}
