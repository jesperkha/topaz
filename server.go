package topaz

import (
	"fmt"
	"net/http"
	"strings"
)

type server struct {
}

func NewServer() Server {
	return &server{}
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

// Returns slice of paramter names in order
func getUrlParams(url string) urlParams {
	split := strings.Split(url, "/")[1:]
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

func (s *server) handle(method string, path string, handlerFunc Handler) {
	params := getUrlParams(path)

	// Todo: handle cases for queries in url
	httpHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// Todo: make coupled function for creating req and res
		res := response{response: w, request: r}
		req := request{
			request: r,
			params:  map[string]string{},
			query:   map[string]string{},
		}

		urlSplit := strings.Split(r.URL.Path[1:], "/")
		for _, param := range params.parameters {
			// Index cannot be out of range if identifiers match in length
			req.query[param.name] = urlSplit[param.index]
			urlSplit[param.index] = "-"
		}

		// Non-matching url identifier patterns
		if fmt.Sprintf("/%s", strings.Join(urlSplit, "/")) != params.identifier {
			return
		}

		handlerFunc(&req, &res)
		if res.status == 0 {
			w.WriteHeader(http.StatusOK)
		}
	}

	// Get leading path before any variable url parameters
	rootPath := "/"
	if split := strings.Split(path, "/")[1:]; len(split) > 1 {
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

func (s *server) Get(path string, handlerFunc Handler) {
	s.handle(http.MethodGet, path, handlerFunc)
}

func (s *server) Post(path string, handlerFunc Handler) {
	s.handle(http.MethodPost, path, handlerFunc)
}

func (s *server) Static(entryPoint string) error {

	return nil
}

func (s *server) Listen(port string) error {
	return http.ListenAndServe(port, nil)
}

func (s *server) Close() {
}
